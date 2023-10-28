package strategies

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/voyagesez/auservice/src/configs"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/utils"
	"golang.org/x/oauth2"
)

type OAuthStrategies struct{}
type GoogleStrategy struct{}
type FacebookStrategy struct{}
type GithubStrategy struct{}

type OAuthProfile struct {
	Sub       string `json:"sub"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type FacebookOAuthProfile struct{}
type GoogleOAuthProfile struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"picture"`
}
type GithubOAuthProfile struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type OAuthStrategiesImpl interface {
	Handler(r *http.Request) (*OAuthProfile, error)
}

func (s *OAuthStrategies) GetOauthStrategy(provider string) OAuthStrategiesImpl {
	switch provider {
	case constants.Google:
		return &GoogleStrategy{}
	case constants.Facebook:
		return &FacebookStrategy{}
	case constants.Github:
		return &GithubStrategy{}
	default:
		return nil
	}
}

func (g *GoogleStrategy) Handler(r *http.Request) (*OAuthProfile, error) {
	const googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
	resp, err := getOAuthProfile(r, configs.GetOauth2Configs().Google, googleUserInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile GoogleOAuthProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &OAuthProfile{
		Sub:       fmt.Sprintf("%s|%s", constants.Google, profile.ID),
		Email:     profile.Email,
		Name:      profile.Name,
		AvatarURL: profile.AvatarURL,
	}, nil
}

func (g *GithubStrategy) Handler(r *http.Request) (*OAuthProfile, error) {
	const githubUserInfoURL = "https://api.github.com/user"

	resp, err := getOAuthProfile(r, configs.GetOauth2Configs().Github, githubUserInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile GithubOAuthProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &OAuthProfile{
		Sub:       fmt.Sprintf("%s|%d", constants.Github, profile.ID),
		Email:     profile.Email,
		Name:      profile.Name,
		AvatarURL: profile.AvatarURL,
	}, nil
}

func (f *FacebookStrategy) Handler(r *http.Request) (*OAuthProfile, error) {
	const facebookUserInfoURL = "https://graph.facebook.com/me?fields=email,name,picture"
	resp, err := getOAuthProfile(r, configs.GetOauth2Configs().Facebook, facebookUserInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return nil, nil
}

func getOAuthProfile(r *http.Request, conf *oauth2.Config, url string) (*http.Response, error) {
	ctx := r.Context()
	code := r.FormValue("code")

	if utils.IsEmptyString(code) {
		log.Println("func getOAuthProfile__error: code is empty")
		return nil, errors.New("code is empty")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Printf("func getOAuthProfile__error: %s\n", err.Error())
		return nil, err
	}
	client := conf.Client(ctx, token)
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("func getOAuthProfile__error: %s\n", err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("func getOAuthProfile__error: status code %d\n", resp.StatusCode)
		return nil, errors.New("failed to get user profile")
	}
	return resp, nil
}
