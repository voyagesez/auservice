package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func ValidateAccessToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		_, ok := claims["email"].(string)
		if !ok {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"message": "token is invalid or expired",
				"error":   "unauthorized",
			})
			return
		}
		h.ServeHTTP(w, r)
	})
}
