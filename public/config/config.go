package config

import (
	"os"
	"sync"

	"github.com/eryajf/eryajfctl/public/logger"
	"gopkg.in/yaml.v3"
)

// Configuration 项目配置
type Configuration struct {
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		// 默认从~/.config.yml 读取，如果读不到，则从当前目录下 config.yml 读取
		data, err := os.ReadFile(os.ExpandEnv("$HOME/.config.yml"))
		if err != nil {
			data, err = os.ReadFile("config.yml")
			if err != nil {
				logger.Fatal(err)
			}
		}
		// 此处如果使用go-bindata将本地配置文件转成二进制，那么使用如下代码读取配置文件
		// data, err := Asset("config.yml")
		if err != nil {
			logger.Fatal(err)
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			logger.Fatal(err)
		}

		// 如果环境变量有配置，读取环境变量
		username := os.Getenv("USERNAME")
		if username != "" {
			config.UserName = username
		}
		password := os.Getenv("PASSWORD")
		if password != "" {
			config.PassWord = password
		}
	})

	return config
}
