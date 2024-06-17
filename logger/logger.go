package logger

import (
	"github.com/google/wire"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
)

var (
	ProviderSet = wire.NewSet(NewLogger)
)

type LoggerConfig struct {
	Type string
}

func NewLogger(config *config.Config) logf.Logger {
	logConf := new(LoggerConfig)
	err := config.UnmarshalKey("logger", logConf)
	if err != nil {
		panic(err)
	}
	switch logConf.Type {
	case "zap":
		return logf.NewLoggerWithZap()
	case "logrus":
		return logf.NewLoggerWithLogrus()
	case "zerolog":
		return logf.NewLoggerWithZeroLog()
	default:
		return logf.NewLoggerWithZap()
	}
}
