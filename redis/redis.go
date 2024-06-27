package redis

import (
	"github.com/google/wire"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewClient)

func NewClient(config *config.Config, log logf.Logger) (redis.Cmdable, error) {
	redisType := config.GetString("redis.type")
	switch redisType {
	case "single":
		return NewSingleClient(config, log)
	case "cluster":
		return NewClusterClient(config, log)
	default:
		return NewSingleClient(config, log)
	}
}

func NewSingleClient(config *config.Config, log logf.Logger) (*redis.Client, error) {
	opt := new(redis.Options)
	err := config.UnmarshalKey("redis", opt)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return redis.NewClient(opt), nil
}

func NewClusterClient(config *config.Config, log logf.Logger) (*redis.ClusterClient, error) {
	opt := new(redis.ClusterOptions)
	err := config.UnmarshalKey("redis", opt)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return redis.NewClusterClient(opt), nil
}
