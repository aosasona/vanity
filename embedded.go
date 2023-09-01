package embedded

import (
	"embed"
)

//go:embed config.json
var Config []byte

//go:embed schema.json
var Schema []byte

//go:embed web/templates/*
var Templates embed.FS
