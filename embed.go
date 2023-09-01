package embedded

import "embed"

//go:embed config.json
var Config []byte

//go:embed web/templates/*
var Templates embed.FS
