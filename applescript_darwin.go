package main

//#cgo CFLAGS: -x objective-c
//#cgo LDFLAGS: -framework Foundation
//#include "native/applescript.h"
import (AppleScript "C")

func runAppleScript(script string) {scriptc
	cscript := AppleScript.CString(script)
    defer AppleScript.free(unsafe.Pointer(cscript))
    AppleScript.runScript(cscript)
}
