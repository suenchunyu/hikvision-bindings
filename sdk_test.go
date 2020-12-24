package hik_vision_sdk

import "testing"

func TestInit(t *testing.T) {
	rec := make(chan Package)
	defer close(rec)
	sdk := Init(&HikVisionSDKConfig{
		MaxLoginUser: 2048,
		MaxAlarmNum:  2048,
		SDKPath:      "",
		SourceAddr:   "hik://localhost:8000|root:root",
		LinkMode:     0,
		StreamType:   0,
	})
	defer sdk.Release()
	if err := sdk.OpenRealTimePlayer(); err != nil {
		panic(err)
	}
	sdk.RegistryReceiver(&rec)
	for {
		select {
		case pkg := <-rec:
			t.Logf("Received package, data length: %d", len(pkg.Data))
		default:
			t.Log("Not received!")
		}
	}
}
