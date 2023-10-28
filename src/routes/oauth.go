package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/voyagesez/auservice/src/configs"
	"github.com/voyagesez/auservice/src/internals/handlers"
)

func NewOauthRoutes(r chi.Router) {
	oauthConfigs := configs.GetOauth2Configs()
	oauthHandlers := handlers.NewOAuthHandlers(&oauthConfigs)

	r.Route("/oauth", func(r chi.Router) {
		r.Get("/{provider}/authorize", oauthHandlers.Authorize)
		r.Get("/{provider}/token", oauthHandlers.Token)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/refresh-token", oauthHandlers.RefreshToken)
		r.Post("/logout", oauthHandlers.Logout)
	})

}
