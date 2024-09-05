package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.trulyao.dev/vanity/config"
	"go.trulyao.dev/vanity/web"
)

func main() {
	initConfig := flag.Bool("init", false, "Initialize a new configuration file")
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if *initConfig {
		if err := config.CreateDefaultConfig(*configPath); err != nil {
			fmt.Printf("Failed to create default configuration: %s", err.Error())
		}
		return
	}

	r := chi.NewRouter()

	config, err := config.Load(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	web.SetConfig(&config)

	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", web.Index)
	r.Get("/schemas/{schema}", web.Schemas)
	r.Get("/{package}", web.ServePackage)
	r.Get("/{package}/cmd/*", web.ServePackage)
	r.Get("/{package}/v{version}", web.ServePackage)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,
	}

	log.Info().Msgf("Starting server on port %d", config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
