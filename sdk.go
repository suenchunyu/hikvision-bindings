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
	env *HikVisionEnv
}

var DefaultConfig = &HikVisionSDKConfig{
	LinkMode:   int(TCP),
	StreamType: int(MainStream),
}

func Init(config *HikVisionSDKConfig) *HikVisionSDK {
	env := initEnv(config)
	return &HikVisionSDK{
		env: env,
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

// TODO: Get real time player's stream data from callback function
