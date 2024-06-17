//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package example

import (
	"github.com/google/wire"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/oldsaltedfish/wireprovider/kafka"
	"github.com/oldsaltedfish/wireprovider/logger"
)

// wireApp init kratos application.
//

// wireApp init kratos application.
func wireApp() (*App, error) {
	panic(wire.Build(logger.ProviderSet, kafka.ProviderSet, config.ProviderSet, ProviderSet))
}
