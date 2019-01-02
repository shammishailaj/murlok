// +build darwin

package mac

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework WebKit
#include "app.h"
*/
import "C"
import (
	"unsafe"

	"github.com/maxence-charriere/murlok/internal/core"
)

var (
	platform = core.Platform{
		Handler: platformCall,
	}

	golang core.Go
)

func platformCall(call string) error {
	ccall := C.CString(call)
	C.platformCall(ccall)
	C.free(unsafe.Pointer(ccall))
	return nil
}

//export platformReturn
func platformReturn(returnID, out, err *C.char) {
	platform.Return(
		C.GoString(returnID),
		C.GoString(out),
		C.GoString(err),
	)
}

//export goCall
func goCall(ccall *C.char) {
	call := C.GoString(ccall)

	err := golang.Call(call)
	if err != nil {
		panic(err)
	}
}
