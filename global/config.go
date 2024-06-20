package global

import (
	"github.com/spf13/viper"
	"github.com/wxr343/logger/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	//Log    *zap.Logger
}

var App = new(Application)
