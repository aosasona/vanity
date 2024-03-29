package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Domain   string   `json:"domain"`
	Port     uint     `json:"port"`
	Packages Packages `json:"packages"`
}

func Load(path string) (Config, error) {
	var (
		c   Config
		err error = nil
	)

	configBytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	if err = json.Unmarshal(configBytes, &c); err != nil {
		return Config{}, err
	}

	return c, err
}

func (c Config) String() string {
	return fmt.Sprintf("Domain: %s\nPackages:\n%s", c.Domain, c.Packages)
}
