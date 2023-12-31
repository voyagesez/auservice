package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/voyagesez/auservice/src/configs"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/internals/db"
	"github.com/voyagesez/auservice/src/internals/dtos"
	"github.com/voyagesez/auservice/src/internals/strategies"
	"github.com/voyagesez/auservice/src/internals/usecase"
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
	dbInstance   *db.DatabaseInstance
	authUseCase  usecase.AuthUseCase
}

var oauthStrategies = strategies.OAuthStrategies{}

func NewOAuthHandlers(oauthConfigs *configs.Oauth2Configs, dbInstance *db.DatabaseInstance,
	authUseCase usecase.AuthUseCase,
) OAuthHandler {
	return &OAuthHandlerImpl{
		oauthConfigs: oauthConfigs,
		dbInstance:   dbInstance,
		authUseCase:  authUseCase,
	}
}

func (o *OAuthHandlerImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	clientOrigin := os.Getenv("CLIENT_ORIGIN")

	if utils.IsEmptyString(provider) {
		http.Redirect(w, r, clientOrigin+"/auth/error", http.StatusTemporaryRedirect)
		return
	}

	state := utils.RandomString(32)
	o.dbInstance.RedisClient.Set(r.Context(), "oauth:state:"+state, state, time.Minute*5) // ttl 5 minutes
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
	ctx := r.Context()
	provider := chi.URLParam(r, "provider")
	code := r.FormValue("code")
	state := r.FormValue("state")

	if utils.IsEmptyString(provider) || utils.IsEmptyString(code) || utils.IsEmptyString(state) {
		errorResponse(w, r, http.StatusBadRequest, ErrorResponse{
			Message: "provider, code and state are required",
			Error:   constants.BadRequest,
		})
		return
	}

	if stateInRedis := o.dbInstance.RedisClient.Get(ctx, "oauth:state:"+state).Val(); utils.IsEmptyString(stateInRedis) {
		errorResponse(w, r, http.StatusUnauthorized, ErrorResponse{
			Message: "state not found",
			Error:   constants.Unauthorized,
		})
		return
	}

	o.dbInstance.RedisClient.Del(ctx, "oauth:state:"+state)
	strategy := oauthStrategies.GetOauthStrategy(provider)
	if strategy == nil {
		errorResponse(w, r, http.StatusBadRequest, ErrorResponse{
			Message: "we don't support this provider",
			Error:   constants.InvalidProvider,
		})
		return
	}

	profile, err := strategy.Handler(r)
	if err != nil {
		errorResponse(w, r, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Error:   constants.BadRequest,
		})
		return
	}

	o.authUseCase.ExternalLogin(ctx, profile, o.authUseCase.ExternalRegister)

	token, err := utils.GenerateAccessToken(dtos.UserAccountResponse{
		UID:      pgtype.UUID{},
		Email:    profile.Email,
		FullName: profile.Name,
		Avatar:   profile.AvatarURL,
	})
	if err != nil {
		errorResponse(w, r, http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
			Error:   constants.InternalServer,
		})
		return
	}

	successResponse(w, r, http.StatusOK, SuccessResponse{
		Message: "success",
		Data:    token,
	})
}

func (o *OAuthHandlerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func (o *OAuthHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {

}
