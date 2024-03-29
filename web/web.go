package web

import "go.trulyao.dev/vanity/config"

var vanityConfig *config.Config

func SetConfig(config *config.Config) {
	vanityConfig = config
}
