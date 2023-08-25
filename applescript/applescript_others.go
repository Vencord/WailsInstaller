//go:build !darwin

package applescript

func runAppleScript(script string) {
	panic("AppleScript on non-Darwin")
}
