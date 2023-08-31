package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type Package struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}

type Packages []Package

//go:embed packages.json
var pBytes []byte

var packages Packages

func init() {
	p := new(struct {
		Packages Packages `json:"packages"`
	})

	if err := json.Unmarshal(pBytes, p); err != nil {
		panic(err)
	}

	packages = p.Packages
}

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
		s += fmt.Sprintf("%s: %s\n", pkg.Name, pkg.Target)
	}

	return s
}
