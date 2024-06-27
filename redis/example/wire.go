//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package example

import (
	"github.com/google/wire"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/oldsaltedfish/wireprovider/logger"
	"github.com/oldsaltedfish/wireprovider/redis"
	goredis "github.com/redis/go-redis/v9"
)

func wireRedis() (goredis.Cmdable, error) {
	panic(wire.Build(redis.ProviderSet, config.ProviderSet, logger.ProviderSet))
}
