package gorm

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ProviderSet = wire.NewSet(New)
)

func New(sqlDB *sql.DB, config *config.Config, log logf.Logger) (db *gorm.DB, err error) {
	gormConfig := &gorm.Config{}
	err = config.UnmarshalKey("gorm", gormConfig)
	if err != nil {
		log.Errorf("config.UnmarshalKey error %v", err)
		return
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), gormConfig)
	if err != nil {
		log.Errorf("gorm.Open error %v", err)
		return
	}
	return
}
