package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/internals/db"
	"github.com/voyagesez/auservice/src/routes"
)

func main() {
	r := chi.NewRouter()
	r.Route("/u", routes.NewOauthRoutes)
	r.Route("/api/v1", routes.NewAPIsRoutes)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	go func() {
		if err := s.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			log.Fatal(`Server start failed: `, err.Error())
		}
	}()
	log.Println("Server started")

	// graceful shutdown
	quit := make(chan os.Signal, 1) // 1 is buffer size
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(`Server shutdown failed: `, err.Error())
	}

	log.Println("Server gracefully stopped")

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(`load env failed: `, err.Error())
	}
	constants.JWTAuthenticator = jwtauth.New("HS256", []byte("access_token_secrets"), nil)
	db.ConnectDatabase()
}
