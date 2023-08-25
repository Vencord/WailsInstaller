// Vencord ShiggyClicker, a little shiggy game :D
// Copyright (c) 2023 Vendicated, Lewis Crichton, and Vencord contributors
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Shiggy Clicker",
		Width:         1024,
		Height:        768,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        runtime.GOOS == "windows",
		BackgroundColour: &options.RGBA{R: 40, G: 40, B: 40, A: 1},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				UseToolbar:                 false,
				FullSizeContent:            true,
			},
			Appearance: mac.NSAppearanceNameDarkAqua,
			About: &mac.AboutInfo{
				Title: "Shiggy Clicker",
				Message: "A cross-platform shiggy game.",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}