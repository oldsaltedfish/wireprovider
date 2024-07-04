// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package example

import (
	"database/sql"
	"github.com/oldsaltedfish/wireprovider/config"
	"github.com/oldsaltedfish/wireprovider/database"
	"github.com/oldsaltedfish/wireprovider/logger"
)

// Injectors from wire.go:

func wireDB() (*sql.DB, error) {
	configConfig := config.NewConfig()
	logfLogger := logger.NewLogger(configConfig)
	db, err := database.New(configConfig, logfLogger)
	if err != nil {
		return nil, err
	}
	return db, nil
}