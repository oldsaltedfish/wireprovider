package config

import (
	"flag"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"strings"
)

var (
	ProviderSet = wire.NewSet(NewConfig)
	configPath  string
)

type Config struct {
	*viper.Viper
}

func init() {
	flag.StringVar(&configPath, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}
func NewConfig() *Config {
	c := &Config{
		viper.New(),
	}
	lastSlashIndex := strings.LastIndex(configPath, "/")
	var path, configName string

	if lastSlashIndex == -1 {
		path = "./"
		configName = configPath
	} else {
		path = configPath[:lastSlashIndex]
		configName = configPath[lastSlashIndex+1:]
	}
	c.AddConfigPath(path)
	lastPointIndex := strings.LastIndex(configName, ".")
	if lastPointIndex == -1 {
		c.SetConfigName(configName)
	} else {
		c.SetConfigName(configName[:lastPointIndex])
		c.SetConfigType(configPath[lastPointIndex+1:])
	}
	err := c.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return c
}
