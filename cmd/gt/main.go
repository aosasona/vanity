package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.trulyao/config"
	"go.trulyao/web"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", web.Index)
	r.Get("/{package}", web.ServePackage)
	r.Get("/schemas/{schema}", web.Schemas)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port()),
		Handler: r,
	}

	log.Info().Msgf("Starting server on port %d", config.Port())
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
