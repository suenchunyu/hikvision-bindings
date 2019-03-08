package hik_vision_sdk

//#cgo CFLAGS: -I./include/hikvision -I./include/internal
//#cgo LDFLAGS: -L./lib64/hikvision -lhcnetsdk
/*

void CALLBACK stdRealTimeDataCallBack(LONG lRealHandle, DWORD dwDataType, BYTE *pBuffer, DWORD dwBufSize, DWORD dwUser) {
	switch(dwDataType) {
		// ONLY STREAM DATA
		case NET_DVR_STREAMDATA:

			Package pkg = (Package*) malloc(sizeof(Package));

			pkg->data = pBuffer;
			pkg->length = dwBufSize;

			publishPackage(pkg);
			break;
		default:
			// DO NOTHING...
	}
}

#include <stdlib.h>
#include "HCNetSDK.h"
#include "chan.h"
*/
import "C"

var blobChan chan Package

func init() {
	blobChan = make(chan Package)
}

func openPlayer(env *HikVisionEnv) {

}

func closePlayer(env *HikVisionEnv) {

}
