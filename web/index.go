package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	embedded "go.trulyao"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// TODO: implement the index someday
	w.Write([]byte("Nothing on this page..."))
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
		schemaBytes = embedded.Schema
	default:
		http.Error(w, "Invalid schema", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(schemaBytes)
	return
}
