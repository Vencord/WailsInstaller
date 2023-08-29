/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package installer

import (
	"fmt"
	path "path/filepath"
	"strings"

	"os"

	"syscall"
)

var macosNames = map[string]string{
	"stable": "Discord.app",
	"ptb":    "Discord PTB.app",
	"canary": "Discord Canary.app",
	"dev":    "Discord Development.app",
}

func ParseDiscord(p, branch string) *DiscordInstall {
	if !ExistsFile(p) {
		return nil
	}

	resources := path.Join(p, "/Contents/Resources")
	if !ExistsFile(resources) {
		return nil
	}

	if branch == "" {
		branch = GetBranch(strings.TrimSuffix(p, ".app"))
	}

	app := path.Join(resources, "app")
	return &DiscordInstall{
		path:             p,
		branch:           branch,
		appPath:          app,
		isPatched:        ExistsFile(app) || IsDirectory(path.Join(resources, "app.asar")),
		isFlatpak:        false,
		isSystemElectron: false,
	}
}

func FindDiscords() []any {
	var discords []any
	for branch, dirname := range macosNames {
		p := "/Applications/" + dirname
		if discord := ParseDiscord(p, branch); discord != nil {
			fmt.Println("Found Discord Install at", p)
			discords = append(discords, discord)
		}
	}
	return discords
}

func PreparePatch(di *DiscordInstall) {}

func FixOwnership(_ string) error {
	return nil
}

func CheckScuffedInstall() bool {
	return false
}
func CheckForOwnership(path string) bool {

	// Check if file is owned by root

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error while checking for root:", err)
		return false
	}

	uid := fileInfo.Sys().(*syscall.Stat_t).Uid

	gid := fileInfo.Sys().(*syscall.Stat_t).Gid

	if gid != 0 {
		return false
	}

	currentUid := os.Getuid()

	return currentUid != int(uid)
}
