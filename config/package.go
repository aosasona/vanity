package config

import "fmt"

type (
	PackageType string

	Repository struct {
		Host  string `json:"host"`
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Package struct {
		Name    string      `json:"name"`
		Repo    Repository  `json:"repo"`
		Type    PackageType `json:"type"`
		SubPath string      `json:"subPath"`
		Readme  string      `json:"readme"`
	}

	Packages []Package
)

const (
	Module     PackageType = "module"
	Executable PackageType = "executable"
	Project    PackageType = "project"
)

func (p Packages) Len() int           { return len(p) }
func (p Packages) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p Packages) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Packages) Get(name string) (Package, bool) {
	for _, pkg := range p {
		if pkg.Name == name {
			return pkg, true
		}
	}

	return Package{}, false
}

func (p Packages) String() string {
	var s string

	for _, pkg := range p {
		s += fmt.Sprintf("\t-%s: %s/%s/%s\n", pkg.Name, pkg.Repo.Host, pkg.Repo.Owner, pkg.Repo.Name)
	}

	return s
}
