//go:build !darwin

package main

func runAppleScript(script string) {
	panic("AppleScript on non-Darwin")
}
