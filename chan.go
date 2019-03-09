package hik_vision_sdk

//#cgo CFLAGS: -I./internal
/*
#include <stdlib.h>
#include "chan.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

//export publishPackage
func publishPackage(p *C.Package) {
	var blob []byte = C.GoBytes(p.data, p.length)
	defer C.free(unsafe.Pointer(p))

	log.Printf("Received 1 package, length: %d\n", len(blob))
	blobChan, ok := blobChanMap.Load(int(p.id))
	if !ok {
		panic("Cannot load current environment's blob channel")
	}
	*blobChan.(*chan Package) <- Package{
		Data: blob,
	}
}
