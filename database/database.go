package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
	"time"
)

var (
	ProviderSet = wire.NewSet(New)
)

type Config struct {
	DriverName      string
	DataSourceName  string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

const configPath = "database"

func New(config *config.Config, log logf.Logger) (db *sql.DB, err error) {
	dbConfig := new(Config)
	err = config.UnmarshalKey(configPath, dbConfig)
	if err != nil {
		log.Errorf("database config unmarshal error: %v", err)
		return
	}
	db, err = sql.Open(dbConfig.DriverName, dbConfig.DataSourceName)
	if err != nil {
		log.Errorf("open %s error: %v", dbConfig.DriverName, err)
		return
	}
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxIdleTime(dbConfig.ConnMaxIdleTime)
	return
}
