package config

var Config *Configuration

func InitSvc() {
	// 加载配置
	Config = LoadConfig()
}
