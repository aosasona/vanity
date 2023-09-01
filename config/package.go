package config

import "fmt"

type (
	Repository struct {
		Host  string `json:"host"`
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Package struct {
		Name string     `json:"name"`
		Repo Repository `json:"repo"`
	}

	Packages []Package
)

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
