// Vencord WailsInstaller, a cross-platform installer for Vencord.
// Copyright (c) 2023 Vendicated, Lewis Crichton, and Vencord contributors
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Vencord Installer",
		Width:  1024,
		Height: 768,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless: runtime.GOOS == "windows",
		BackgroundColour: &options.RGBA{R: 40, G: 40, B: 40, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
