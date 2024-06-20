package bootstrap

import (
	"github.com/spf13/viper"
	"github.com/wxr343/logger/global"
)

func InitializeConfig() *viper.Viper {
	// 默认config路径
	//config := "../logConfig-example.yaml"
	//// 通过环境变量配置的路径
	//if configEev := os.Getenv("LOGGER_CONFIG"); configEev != "" {
	//
	//	config = configEev
	//
	//	// 初始化viper
	//	v := viper.New()
	//	// config路径
	//	v.SetConfigFile(config)
	//	// config类型
	//	v.SetConfigType("yaml")
	//	if err := v.ReadInConfig(); err != nil {
	//		panic(fmt.Errorf("read config failed: %s \n", err))
	//	}
	//	// 开始监听配置
	//	v.WatchConfig()
	//	v.OnConfigChange(func(in fsnotify.Event) {
	//		log.Println("config file changed:", in.Name)
	//		// 重载配置
	//		if err := v.Unmarshal(&global.App.Config); err != nil {
	//			panic(err)
	//		}
	//	})
	//	// 将配置赋值给全局变量
	//	if err := v.Unmarshal(&global.App.Config); err != nil {
	//		panic(err)
	//	}
	//	return v
	//} else
	{ //默认配置
		global.App.Config.Log.ShowLine = true
		global.App.Config.Log.Format = "json"
		global.App.Config.Log.Level = "info"
		global.App.Config.Log.MaxBackups = 3
		global.App.Config.Log.MaxSize = 500
		global.App.Config.Log.RootDir = "./logs"
		global.App.Config.Log.MaxAge = 28
		global.App.Config.Log.Compress = true
		global.App.Config.Log.Filename = []string{"app_log.json"}
	}
	return nil
}
