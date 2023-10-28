package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/voyagesez/auservice/src/internals/handlers"
)

func NewUserRoutes(r chi.Router) {
	userHandlers := handlers.NewUserHandlers()

	r.Get("/", userHandlers.GetMyProfile)
	r.Get("/{slug}", userHandlers.GetAnotherProfile)
}
