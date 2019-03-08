package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*
#include <stdlib.h>
#include "HCNetSDK.h"
*/
import "C"
import (
	"strconv"
	"strings"
	"unsafe"
)

const (
	CFGSetAbility = 1
	CFGSetSDKPath = 2
	SUCCEED       = 1
	LoginFailed   = -1
)

type HikVisionEnv struct {
	DeviceInfo *C.NET_DVR_DEVICEINFO
	UserID     int
	Config     *HikVisionSDKConfig
	PlayerHdl  int
}

func initEnv(config *HikVisionSDKConfig) *HikVisionEnv {
	ip, port, username, password := parseSourceAddr(config.SourceAddr)
	var device C.NET_DVR_DEVICEINFO
	cIp := C.CString(ip)
	defer C.free(unsafe.Pointer(cIp))
	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))
	cPassword := C.String(password)
	defer C.free(unsafe.Pointer(cPassword))
	var userId int

	// set ability
	var abilityCfg = new(SDKCfgAbility)
	abilityCfg.EnumMaxLoginUsersNum = uint32(config.MaxLoginUser)
	abilityCfg.EnumMaxAlarmNum = uint32(config.MaxAlarmNum)

	// set sdk path

	// init sdk
	C.NET_DVR_SetSDKInitCfg(CFGSetAbility, unsafe.Pointer(abilityCfg))
	if rt := ErrorCode(C.NET_DVR_Init()); rt != SUCCEED {
		goto Error
	}

	// active device
	if userId = int(C.NET_DVR_Login(cIp, C.WORD(port), cUsername, cPassword,
		(*C.NET_DVR_DEVICEINFO)(unsafe.Pointer(&device)))); userId == LoginFailed {
		goto Error
	}

Error:
	errCode := getLastError()
	panic(getErrorMessage(errCode))

	return &HikVisionEnv{
		DeviceInfo: &device,
		UserID:     userId,
		Config:     config,
	}
}

func parseSourceAddr(url string) (ip string, port int, username string, password string) {
	urls := strings.Split(url, "@")
	serverAndPort := strings.Split(urls[0], ":")
	usernameAndPassword := strings.Split(urls[1], ":")
	ip = serverAndPort[0]
	port, err := strconv.Atoi(serverAndPort[1])
	if err != nil {
		panic("Parse source address failure")
	}
	username = usernameAndPassword[0]
	password = usernameAndPassword[1]
	return
}

func (e *HikVisionEnv) release() {
	if rt := int(C.NET_DVR_Logout(C.LONG(e.UserID))); rt != SUCCEED {
		goto Error
	}
	if rt := int(C.NET_DVR_Cleanup()); rt != SUCCEED {
		goto Error
	}
Error:
	errCode := getLastError()
	panic(getErrorMessage(errCode))
}
