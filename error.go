package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*
#include <stdlib.h>
#include "HCNetSDK.h"
*/
import "C"

type ErrorCode int

var errorsMap = map[ErrorCode]string{
	0:  "No error",
	1:  "Invalid username or password",
	2:  "Not enough permissions",
	3:  "SDK not init",
	4:  "Invalid channel code",
	5:  "Out of max link number",
	6:  "Mismatch device and SDK version",
	7:  "Failed to connect to device",
	8:  "Send to device failure",
	9:  "Receive data from device failure",
	10: "Receive data from device timeout",
	11: "Invalid data to send device",
	12: "Invalid call order",
	13: "Current user doesn't have enough permissions",
	14: "Device command timeout",
}

func getErrorMessage(n ErrorCode) string {
	return errorsMap[n]
}

func getLastError() ErrorCode {
	return ErrorCode(C.NET_DVR_GetLastError())
}
