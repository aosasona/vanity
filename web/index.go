package web

import (
	"html/template"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"go.trulyao.dev/vanity/config"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templateString, err := Templates.ReadFile("templates/index.tmpl.html")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := template.New("package").Parse(string(templateString))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	packages := vanityConfig.Packages
	sort.Sort(packages)

	err = t.Execute(w, struct {
		Packages config.Packages
		Host     string
	}{
		Packages: packages,
		Host:     vanityConfig.Domain,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to execute template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Schemas(w http.ResponseWriter, r *http.Request) {
	schema := chi.URLParam(r, "schema")
	if schema == "" {
		http.Error(w, "No schema specified", http.StatusBadRequest)
		return
	}

	var schemaBytes []byte
	switch schema {
	case "config.json":
		schemaBytes = Schema
	default:
		http.Error(w, "Invalid schema", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(schemaBytes)
	return
}
