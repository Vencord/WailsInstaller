// Vencord WailsInstaller, a cross-platform installer for Vencord.
// Copyright (c) 2023 Vendicated, Lewis Crichton, and Vencord contributors
// SPDX-License-Identifier: GPL-3.0-or-later

package installer

import (
	"context"
	"fmt"

	"os/exec"
	"os"
	"syscall"
	"runtime"
	WailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type DiscordData struct {
	Branch     string `json:"branch"`
	Path       string `json:"path"`
	IsPatched  bool   `json:"isPatched"`
	IsOpenAsar bool   `json:"isOpenAsar"`
}

type Installer struct {
	ctx context.Context
}

func NewInstaller() *Installer {
	return &Installer{}
}

func (i *Installer) Startup(ctx context.Context) {
	i.ctx = ctx

	InitGithubDownloader()
	<-GithubDoneChan
}

func (i *Installer) GetLatestVersion() string {
	return LatestHash
}

func (i *Installer) GetInstalledVersion() string {
	return InstalledHash
}

func (i *Installer) ListInstalls() []DiscordData {
	discords := FindDiscords()

	data := make([]DiscordData, len(discords))
	for i, d := range discords {
		in := d.(*DiscordInstall)
		data[i] = DiscordData{
			Path:       in.path,
			Branch:     in.branch,
			IsPatched:  in.isPatched,
			IsOpenAsar: in.IsOpenAsar(),
		}
	}

	return data
}

func (i *Installer) ChooseCustomInstall() (string, error) {
	directory, err := WailsRuntime.OpenDirectoryDialog(i.ctx, WailsRuntime.OpenDialogOptions{
		Title: "Discord Install Directory",
		CanCreateDirectories: false,
		TreatPackagesAsDirectories: true, // for mac, since we need to select the app package
	})

	if err != nil {
		return "", err
	}

	return directory, nil
}

func (i *Installer) Patch(path string) error {
	discordInstall := ParseDiscord(path, "")
	if discordInstall == nil {
		return fmt.Errorf("not a valid Discord install: %s", path)
	}

	return discordInstall.patch()
}

func (i *Installer) Unpatch(path string) error {
	discordInstall := ParseDiscord(path, "")
	if discordInstall == nil {
		return fmt.Errorf("not a valid Discord install: %s", path)
	}

	return discordInstall.unpatch()
}

func (i *Installer) Repair(path string) error {
	discordInstall := ParseDiscord(path, "")
	if discordInstall == nil {
		return fmt.Errorf("not a valid Discord install: %s", path)
	}

	err := InstallLatestBuilds()
	if err != nil {
		return err
	}

	return discordInstall.patch()
}

func (i *Installer) InstallOpenAsar(path string) error {
	discordInstall := ParseDiscord(path, "")
	if discordInstall == nil {
		return fmt.Errorf("not a valid Discord install: %s", path)
	}

	return discordInstall.InstallOpenAsar()
}

func (i *Installer) UninstallOpenAsar(path string) error {
	discordInstall := ParseDiscord(path, "")
	if discordInstall == nil {
		return fmt.Errorf("not a valid Discord install: %s", path)
	}

	return discordInstall.UninstallOpenAsar()
}

func (i *Installer) GetBaseDir() string {
	return BaseDir
}

func InstallLatestBuilds() error {
	return installLatestBuilds()
}

func HandleScuffedInstall() {
	// TODO
}

func (i *Installer) PromptForChown(path string) bool {

	if runtime.GOOS != "darwin" {
		return false
	}

	fmt.Println("Showing osascript prompt", path, "...")

	uid := os.Getuid()
	
	cmd := exec.Command("osascript", "-e", `do shell script "chown -R \"` + fmt.Sprint(uid) + `:wheel\" '` + path + `'" with prompt "Vencord Installer needs to fix Discord file ownership" with administrator privileges`)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error while prompting for chown:", err)
		return false
	}

	return true
}

func (i *Installer) CheckForOwnership(path string) bool {

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