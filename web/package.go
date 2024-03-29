package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"go.trulyao.dev/vanity/config"
)

type PackageVars struct {
	Domain      string
	PackageName string
	RepoURL     string
}

func ServePackage(w http.ResponseWriter, r *http.Request) {
	pkg := chi.URLParam(r, "package")
	if pkg == "" {
		http.Error(w, "No package specified", http.StatusBadRequest)
		return
	}

	templateString, err := Templates.ReadFile("templates/package.tmpl.html")
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

	p, exists := vanityConfig.Packages.Get(strings.ToLower(pkg))
	if !exists {
		log.Error().Str("package", pkg).Msg("Package not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repoURL := fmt.Sprintf("https://%s/%s/%s", p.Repo.Host, p.Repo.Owner, p.Repo.Name)
	if p.Type == config.Project {
		// redirect to repo if it is a project
		http.Redirect(w, r, repoURL, http.StatusFound)
		return
	}

	err = t.Execute(w, PackageVars{
		Domain:      vanityConfig.Domain,
		PackageName: p.Name,
		RepoURL:     repoURL,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to execute template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
