//go:build wireinject
// +build wireinject

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
package example

import (
	"github.com/google/wire"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/oldsaltedfish/wireprovider/database"
	gormProvider "github.com/oldsaltedfish/wireprovider/gorm"
	"github.com/oldsaltedfish/wireprovider/logger"
	"gorm.io/gorm"
)

func wireDB() (*gorm.DB, error) {
	panic(wire.Build(config.ProviderSet, database.ProviderSet, logger.ProviderSet, gormProvider.ProviderSet))
}
