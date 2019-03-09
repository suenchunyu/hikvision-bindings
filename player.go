package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision -I./include/internal
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk -lhpr
/*
#include <stdlib.h>
#include "HCNetSDK.h"
#include "chan.h"

void CALLBACK stdRealTimeDataCallBack(LONG lRealHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, DWORD dwUser) {
	Package* pkg;
	// ONLY STREAM DATA
	if (dwDataType == NET_DVR_STREAMDATA) {
		pkg = (Package*) malloc(sizeof(Package));

		pkg->data = pBuffer;
		pkg->length = dwBufSize;
		pkg->id = dwUser;

		publishPackage(pkg);
	}
}

LONG wrapper_open_real_time_player(DWORD stream_type, DWORD link_mode, LONG user_id) {
	NET_DVR_PREVIEWINFO preview_info = {0};
	preview_info.lChannel = 1;
	preview_info.dwStreamType = stream_type;
	preview_info.dwLinkMode = link_mode;
	preview_info.bBlocked = 1;

	LONG real_player_handle = NET_DVR_RealPlay_V40(user_id, &preview_info, (REALDATACALLBACK) (stdRealTimeDataCallBack), NULL);

	return real_player_handle;
}
*/
import "C"
import (
	"github.com/pkg/errors"
	"sync"
)

var blobChanMap *sync.Map

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
	blobChanMap = new(sync.Map)
}

func openPlayer(env *HikVisionEnv) error {
	playerHdl := int(C.wrapper_open_real_time_player(C.DWORD(env.Config.StreamType),
		C.DWORD(env.Config.LinkMode), C.LONG(env.UserID)))
	if playerHdl < 0 {
		goto Error
	}

	env.PlayerHdl = playerHdl
	return nil

Error:
	defer env.release()
	errCode := getLastError()
	println(errCode)
	return errors.New(getErrorMessage(errCode))
}

func closePlayer(env *HikVisionEnv) {
	if rt := int(C.NET_DVR_StopRealPlay(C.LONG(env.PlayerHdl))); rt != SUCCEED {
		errCode := getLastError()
		panic(getErrorMessage(errCode))
	}
}
