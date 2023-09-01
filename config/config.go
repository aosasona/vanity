package config

import (
	"encoding/json"
	"fmt"

	embedded "go.trulyao"
)

var configBytes = embedded.Config

type Config struct {
	Domain   string   `json:"domain"`
	Port     uint     `json:"port"`
	Packages Packages `json:"packages"`
}

var c Config

func init() {
	if err := json.Unmarshal(configBytes, &c); err != nil {
		panic(err)
	}
}

func Port() uint {
	return c.Port
}

func Domain() string {
	return c.Domain
}

func Pkgs() Packages {
	return c.Packages
}

func String() string {
	return fmt.Sprintf("Domain: %s\nPackages:\n%s", c.Domain, c.Packages)
}
