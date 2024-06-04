//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package example

import (
	"github.com/google/wire"
	wireconfig "github.com/oldsaltedfish/wireprovider/config"
)

// wireApp init kratos application.
func wireConfig() *wireconfig.Config {
	panic(wire.Build(wireconfig.ProviderSet))
}
