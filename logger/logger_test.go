package logger

import (
	"github.com/wxr343/logger/config"
	"testing"
)

func TestInitializeLog(t *testing.T) {
	logconfig := config.Log{
		Level:      "info",
		RootDir:    "./logs",
		Filename:   []string{"test1.json"},
		Format:     "json",
		ShowLine:   true,
		MaxBackups: 3,
		MaxSize:    500,
		MaxAge:     28,
		Compress:   true,
	}
	C := config.Configuration{
		Log: logconfig,
	}

	InitializeConfig()
	LogInfo := InitializeLog(C)
	LogInfo.Info("C log init success!")

	LogError := InitializeLog()
	LogError.Error("default log init success!")
}
