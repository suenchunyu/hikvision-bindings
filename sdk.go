package hik_vision_sdk

import "errors"

// HikVisionSDKConfig is the specification of HikVision NetSDK,
// SourceAddr: hik://{ip}:{port}@{username}:{password},
// Example: hik://localhost:5050@admin:admin
type HikVisionSDKConfig struct {
	MaxLoginUser int    `toml:"max_login_user"`
	MaxAlarmNum  int    `toml:"max_alarm_num"`
	SDKPath      string `toml:"sdk_path"`
	SourceAddr   string `toml:"source_addr"`
	LinkMode     int    `toml:"real_play_link_mode"`
	StreamType   int    `toml:"real_play_stream_type"`
}

type HikVisionSDK struct {
	env   *HikVisionEnv
	rec   *chan Package
	stopC chan int
}

var DefaultConfig = &HikVisionSDKConfig{
	MaxLoginUser: 2048,
	MaxAlarmNum:  2048,
	SDKPath:      "",
	LinkMode:     int(TCP),
	StreamType:   int(MainStream),
}

func Init(config *HikVisionSDKConfig) *HikVisionSDK {
	env := initEnv(config)
	return &HikVisionSDK{
		env:   env,
		stopC: make(chan int, 8),
	}
}

func (sdk *HikVisionSDK) OpenRealTimePlayer() error {
	if sdk.env.UserID < 0 {
		return errors.New("SDK not initialized")
	}
	return openPlayer(sdk.env)
}

func (sdk *HikVisionSDK) CloseRealTimePlayer() error {
	if sdk.env.UserID < 0 {
		return errors.New("SDK not initialized")
	}
	closePlayer(sdk.env)
	return nil
}

// Only call once!!
func (sdk *HikVisionSDK) RegistryReceiver(rec *chan Package) {
	if rec == nil {
		panic("cannot registry receiver chan")
	} else if sdk.rec != nil {
		panic("RegistryReceiver func only can be call once")
	}
	sdk.rec = rec
	blobChan, ok := blobChanMap.Load(sdk.env.UserID)
	if !ok {
		panic("Cannot load current environment's blob channel")
	}
	go func() {
		select {
		case <-sdk.stopC:
			// receive stop signal
			return
		case pkg := <-*blobChan.(*chan Package):
			*rec <- pkg
		}
	}()
}

func (sdk *HikVisionSDK) UnRegistryReceiver() {
	sdk.stopC <- 1
}

func (sdk *HikVisionSDK) Release() {
	sdk.UnRegistryReceiver()
	sdk.env.release()
}
