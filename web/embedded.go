package web

import (
	"embed"
)

//go:embed templates/schema.json
var Schema []byte

//go:embed templates/*
var Templates embed.FS
