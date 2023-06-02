package bootstrap

import (
	"fmt"
	"ginserver/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func InitializeConfig(args ...string) {
	configPath := "config.yaml"
	// 用于测试环境
	for _, arg := range args {
		if !strings.EqualFold(arg, "") {
			configPath = arg
			break
		}
	}
	// 生产环境可通过环境变量改变配置文件路径
	if path := os.Getenv("VIPER_CONFIG"); path != "" {
		configPath = path
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Read config file failed, err: %+v", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
		// 热重载配置
		if err := v.Unmarshal(&global.AppConfig.Config); err != nil {
			log.Printf("Reload config file %s failed, err: %+v", in.Name, err)
		}
	})
	// 解析配置文件
	if err := v.Unmarshal(&global.AppConfig.Config); err != nil {
		panic(fmt.Sprintf("Unmarshal config file failed, err: %+v", err))
	}
}
