package applescript

//#cgo CFLAGS: -x objective-c
//#cgo LDFLAGS: -framework Foundation
//#include "native/applescript.h"
import "C"

func runAppleScript(script string) {
    cscript := C.CString(script)
    defer C.free(unsafe.Pointer(cscript))
    C.runScript(cscript)
}
