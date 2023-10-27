package configs

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

type Oauth2Configs struct {
	Google   *oauth2.Config
	Facebook *oauth2.Config
	Github   *oauth2.Config
}

func GetOauth2Configs() Oauth2Configs {
	return Oauth2Configs{
		Google: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     endpoints.Google,
		},
		Facebook: &oauth2.Config{
			ClientID:     os.Getenv("FACEBOOK_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("FACEBOOK_OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("FACEBOOK_OAUTH_REDIRECT_URL"),
			Scopes:       []string{"email", "public_profile"},
			Endpoint:     endpoints.Facebook,
		},
		Github: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GITHUB_OAUTH_REDIRECT_URL"),
			Scopes:       []string{"read:user", "user:email"},
			Endpoint:     endpoints.GitHub,
		},
	}
}
