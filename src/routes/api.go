package routes

import "github.com/go-chi/chi/v5"

func NewAPIsRoutes(r chi.Router) {
	r.Route("/user", NewUserRoutes)
}
