package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oldsaltedfish/wireprovider/config"
)

type Config struct {
	DriverName     string
	DataSourceName string
}

func New(config config.Config) {
	config.UnmarshalKey("mysql", &config)
	sqlDB, err := sql.Open("mysql", "mydb_dsn")
}
