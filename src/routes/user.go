package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/internals/handlers"
	"github.com/voyagesez/auservice/src/middleware"
)

func NewUserRoutes(r chi.Router) {
	userHandlers := handlers.NewUserHandlers()

	r.Use(jwtauth.Verifier(constants.JWTAuthenticator), middleware.ValidateAccessToken)
	r.Use(jwtauth.Authenticator)

	r.Get("/", userHandlers.GetMyProfile)
	r.Get("/{slug}", userHandlers.GetAnotherProfile)
}
