package hik_vision_sdk

import "testing"

func TestInit(t *testing.T) {
	_ = Init(&HikVisionSDKConfig{
		MaxLoginUser: 2048,
		MaxAlarmNum:  2048,
		SDKPath:      "",
		SourceAddr:   "hik://127.0.0.1:8000@admin:admin",
		LinkMode:     1,
		StreamType:   1,
	})
}
