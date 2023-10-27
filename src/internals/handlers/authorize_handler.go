package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/voyagesez/auservice/src/configs"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/internals/strategies"
	"github.com/voyagesez/auservice/src/utils"
)

type OAuthHandler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type OAuthHandlerImpl struct {
	oauthConfigs *configs.Oauth2Configs
}

var oauthStrategies = strategies.OAuthStrategies{}

func NewOAuthHandlers(oauthConfigs *configs.Oauth2Configs) OAuthHandler {
	return &OAuthHandlerImpl{
		oauthConfigs: oauthConfigs,
	}
}

func (o *OAuthHandlerImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	provider := r.FormValue("provider")
	clientOrigin := os.Getenv("CLIENT_ORIGIN")

	if utils.IsEmptyString(state) || utils.IsEmptyString(provider) {
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
		return
	}

	switch provider {
	case constants.Google:
		http.Redirect(w, r, o.oauthConfigs.Google.AuthCodeURL(state), http.StatusTemporaryRedirect)
	case constants.Facebook:
		http.Redirect(w, r, o.oauthConfigs.Facebook.AuthCodeURL(state), http.StatusTemporaryRedirect)
	case constants.Github:
		http.Redirect(w, r, o.oauthConfigs.Github.AuthCodeURL(state), http.StatusTemporaryRedirect)
	default:
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
	}

}

func (o *OAuthHandlerImpl) Token(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	provider := chi.URLParam(r, "provider")
	clientOrigin := os.Getenv("CLIENT_ORIGIN")

	if utils.IsEmptyString(code) || utils.IsEmptyString(provider) {
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
		return
	}

	strategy := oauthStrategies.GetOauthStrategy(provider)
	if strategy == nil {
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
		return
	}

	profile, err := strategy.Handler(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
		return
	}
	log.Println(profile)
}

func (o *OAuthHandlerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func (o *OAuthHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {

}
