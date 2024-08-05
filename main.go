package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend
var assets embed.FS

func main() {
	var opts = &options.App{
		Title:            "Image Converter",
		Width:            800,
		Height:           465,
		MinWidth:         435,
		MinHeight:        460,
		WindowStartState: options.Maximised,
		OnStartup:        startup,
		OnDomReady:       domReady,
		OnBeforeClose:    beforeClose,
		OnShutdown:       shutdown,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
		},
	}

	options.MergeDefaults(opts)

	err := wails.Run(opts)

	if err != nil {
		log.Fatal(err)
	}
}
