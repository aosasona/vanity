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
	SourceURL   string
	SubPath     string
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
		http.Redirect(w, r, repoURL, http.StatusFound)
		return
	} else if p.Type == config.Executable {
		http.Redirect(w, r, fmt.Sprintf("%s/releases", repoURL), http.StatusFound)
		return
	}

	cacheMaxAge := int64(86400)
	if vanityConfig.MaxCacheAge > 0 {
		cacheMaxAge = vanityConfig.MaxCacheAge
	}
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", cacheMaxAge))

	source := fmt.Sprintf("%s %s/tree/master{/dir} %s/blob/master{/dir}/{file}#L{line}", repoURL, repoURL, repoURL)
	if strings.Contains(repoURL, "gitlab") {
		source = fmt.Sprintf("%v %v/-/tree/master{/dir} %v/-/blob/master{/dir}/{file}#L{line}", repoURL, repoURL, repoURL)
	} else if strings.Contains(repoURL, "bitbucket") {
		source = fmt.Sprintf("%v %v/src/master{/dir} %v/src/master{/dir}/{file}#lines-{line}", repoURL, repoURL, repoURL)
	}

	if p.SubPath != "" {
		repoURL = fmt.Sprintf("%s/%s", source, strings.Trim(p.SubPath, "/"))
	}

	err = t.Execute(w, PackageVars{
		Domain:      vanityConfig.Domain,
		PackageName: p.Name,
		RepoURL:     repoURL,
		SourceURL:   source,
		SubPath:     p.SubPath,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to execute template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
