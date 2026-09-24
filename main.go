package main

import (
	_ "embed"
	"github.com/crowquillx/silo-anime-themes/internal/app"
	"github.com/crowquillx/silo-theme-songs/pkg/pluginapp"
)

//go:embed manifest.json
var manifestJSON []byte
var version string

func main() { pluginapp.Main(manifestJSON, version, "animethemes", true, app.New) }
