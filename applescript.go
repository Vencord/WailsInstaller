//go:build !darwin

package main

type AppleScript struct {}

func NewAppleScript() *AppleScript {
	return &AppleScript{}
}

func (a *AppleScript) RunScript(script string) {
	// handed off to platform-specific implementation
	runAppleScript(script)
}
