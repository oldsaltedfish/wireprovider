//go:build wireinject
// +build wireinject

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
package example

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/oldsaltedfish/wireprovider/database"
	"github.com/oldsaltedfish/wireprovider/logger"
)

func wireDB() (*sql.DB, error) {
	panic(wire.Build(config.ProviderSet, database.ProviderSet, logger.ProviderSet))
}
