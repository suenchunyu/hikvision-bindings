package hik_vision_sdk

//#cgo CFLAGS: -I./internal
/*
#include <stdlib.h>
#include "chan.h"
*/
import "C"
import "unsafe"

//export publishPackage
func publishPackage(p *C.Package) {
	var blob []byte = C.GoBytes(p.data, p.length)
	defer C.free(unsafe.Pointer(p.data))
	defer C.free(unsafe.Pointer(p))

	BlobChan <- Package{
		blob,
	}
}
