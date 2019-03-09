package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*
#include "HCNetSDK.h"
#include <stdlib.h>
#include <string.h>

	LONG wrapper_login(char *ip,WORD port,char *username,char *password) {
		NET_DVR_USER_LOGIN_INFO login_info = {0};
		login_info.bUseAsynLogin = 0;
		strcpy(login_info.sDeviceAddress, ip);
		login_info.wPort = port;
		strcpy(login_info.sUserName, username);
		strcpy(login_info.sPassword, password);

		NET_DVR_DEVICEINFO_V40 device_info = {0};
		return NET_DVR_Login_V40(&login_info, &device_info);
	}
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
	blobChan := make(chan Package)
	ip, port, username, password := parseSourceAddr(config.SourceAddr)
	cIp := C.CString(ip)
	defer C.free(unsafe.Pointer(cIp))
	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))
	cPassword := C.CString(password)
	defer C.free(unsafe.Pointer(cPassword))
	var userId int

	// set ability
	var abilityCfg = new(SDKCfgAbility)
	abilityCfg.EnumMaxLoginUsersNum = uint32(config.MaxLoginUser)
	abilityCfg.EnumMaxAlarmNum = uint32(config.MaxAlarmNum)

	// set sdk path

	// init sdk
	C.NET_DVR_SetSDKInitCfg(CFGSetAbility, unsafe.Pointer(abilityCfg))
	if rt := int(C.NET_DVR_Init()); rt != SUCCEED {
		goto Error
	}

	userId = int(C.wrapper_login(cIp, C.WORD(port), cUsername, cPassword))
	// active device
	if userId == LoginFailed {
		goto Error
	}

	// New channel for blob to blobChanMap
	blobChanMap.Store(userId, &blobChan)

	return &HikVisionEnv{
		UserID: userId,
		Config: config,
	}

Error:
	errCode := getLastError()
	panic(getErrorMessage(errCode))

}

func parseSourceAddr(url string) (ip string, port int, username string, password string) {
	temp := strings.Split(url, "hik://")
	urls := strings.Split(temp[1], "|")
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
	blobChan, ok := blobChanMap.Load(e.UserID)
	if !ok {
		panic("Cannot load current environment's blob channel")
	}
	defer close(*blobChan.(*chan Package))
	if rt := int(C.NET_DVR_Logout(C.LONG(e.UserID))); rt != SUCCEED {
		println(rt)
		goto Error
	}
	if rt := int(C.NET_DVR_Cleanup()); rt != SUCCEED {
		println(rt)
		goto Error
	}

Error:
	errCode := getLastError()
	panic(getErrorMessage(errCode))
}
