package hik_vision_sdk

// HikVisionSDKConfig is the specification of HikVision NetSDK,
// SourceAddr: hik://{ip}:{port}@{username}:{password},
// Example: hik://localhost:5050@admin:admin
type HikVisionSDKConfig struct {
	MaxLoginUser int    `toml:"max_login_user"`
	MaxAlarmNum  int    `toml:"max_alarm_num"`
	SDKPath      string `toml:"sdk_path"`
	SourceAddr   string `toml:"source_addr"`
}

type HikVisionSDK struct {
	env *HikVisionEnv
}

var DefaultConfig = &HikVisionSDKConfig{}

func Init(config *HikVisionSDKConfig) *HikVisionSDK {
	env := initEnv(config)
	return &HikVisionSDK{
		env: env,
	}
}

// TODO: Open real time player
// TODO: Close real time player
// TODO: Get real time player's stream data from callback function
