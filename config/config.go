package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultConfig = `{
  "$schema": "https://go.trulyao.dev/schemas/config.json",
  "port": 8080,
  "domain": "example.com",
  "maxCacheAge": 3600,
  "packages": [
        {
          "name": "foo",
          "repo": {
            "host": "github.com",
            "owner": "user",
            "name": "foo"
          },
          "type": "module",
          "readme": "https://raw.githubusercontent.com/user/foo/master/README.md"
        }
    ]
}
`

type Config struct {
	Domain      string   `json:"domain"`
	Port        uint     `json:"port"`
	MaxCacheAge int64    `json:"maxCacheAge"`
	Packages    Packages `json:"packages"`
}

func CreateDefaultConfig(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(defaultConfig); err != nil {
		return err
	}

	return nil
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
