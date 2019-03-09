package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*
#include <stdlib.h>
#include "HCNetSDK.h"

	char* wrapper_get_error_msg(LONG code) {
		return NET_DVR_GetErrorMsg(&code);
	}
*/
import "C"

type ErrorCode int

func getErrorMessage(n ErrorCode) string {
	return C.GoString(C.wrapper_get_error_msg(C.int(n)))
}

func getLastError() ErrorCode {
	return ErrorCode(C.NET_DVR_GetLastError())
}
