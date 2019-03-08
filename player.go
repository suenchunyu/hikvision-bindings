package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision -I./include/internal
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*
#include <stdlib.h>
#include "HCNetSDK.h"
#include "chan.h"

void stdRealTimeDataCallBack(LONG lRealHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, DWORD dwUser) {
	Package* pkg;
	// ONLY STREAM DATA
	if (dwDataType == NET_DVR_STREAMDATA) {
		pkg = (Package*) malloc(sizeof(Package));

		pkg->data = pBuffer;
		pkg->length = dwBufSize;

		publishPackage(pkg);
	}
}
*/
import "C"
import (
	"github.com/pkg/errors"
	"unsafe"
)

var BlobChan chan Package

type LinkMode int
type StreamType int

const (
	TCP LinkMode = iota
	UDP
	MultiPlay
	RTP
	RTP_RTSP
	RTP_HTTP

	MainStream StreamType = 0
)

func init() {
	BlobChan = make(chan Package)
}

func openPlayer(env *HikVisionEnv) error {
	// build params
	previewInfo := new(C.NET_DVR_PREVIEWINFO)
	defer C.free(unsafe.Pointer(previewInfo))
	previewInfo.hPlayWnd = 0
	previewInfo.lChannel = C.int(1)
	previewInfo.dwStreamType = C.uint(env.Config.StreamType)
	previewInfo.dwLinkMode = C.uint(env.Config.LinkMode)
	previewInfo.bBlocked = 1

	playerHdl := int(C.NET_DVR_RealPlay_V40(C.int(env.UserID), previewInfo, C.REALDATACALLBACK(C.stdRealTimeDataCallBack), nil))
	if playerHdl < 0 {
		goto Error
	}

	env.PlayerHdl = playerHdl
	return nil

Error:
	defer env.release()
	errCode := getLastError()
	return errors.New(getErrorMessage(errCode))
}

func closePlayer(env *HikVisionEnv) {
	if rt := int(C.NET_DVR_StopRealPlay(C.LONG(env.PlayerHdl))); rt != SUCCEED {
		errCode := getLastError()
		panic(getErrorMessage(errCode))
	}
}
