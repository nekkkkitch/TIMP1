package main

import (
	"embed"
	"log/slog"
	app "window/services/app"
	"window/services/tabler"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the appInst structure

	tabler, err := tabler.NewTabler("table")
	if err != nil {
		slog.Error("Main: no table name given")
	}

	appInst := app.NewApp(tabler)
	// Create application with options
	err = wails.Run(&options.App{
		Title:  "window",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appInst.Startup,
		Bind: []interface{}{
			appInst,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
