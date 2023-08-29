// Vencord WailsInstaller, a cross-platform installer for Vencord.
// Copyright (c) 2023 Vendicated, Lewis Crichton, and Vencord contributors
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"embed"
	"runtime"

	"github.com/vencord/wailsinstaller/installer"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the installer binding
	installer := installer.NewInstaller()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Vencord Installer",
		Width:         1024,
		Height:        768,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        runtime.GOOS == "windows",
		BackgroundColour: &options.RGBA{R: 40, G: 40, B: 40, A: 1},
		OnStartup:        installer.Startup,
		Bind: []any{
			installer,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				UseToolbar:                 false,
				FullSizeContent:            true,
			},
			Appearance: mac.NSAppearanceNameDarkAqua,
			About: &mac.AboutInfo{
				Title:   "Vencord Installer",
				Message: "A cross-platform installer for Vencord.",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
