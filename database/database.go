package database

import (
	"database/sql"
	"time"

	"github.com/indraprasetya154/golang-restful-api/helper"
	"github.com/spf13/viper"
)

func NewDB() *sql.DB {
	db, err := sql.Open(viper.GetString("DB_DRIVER"), viper.GetString("DB_USERNAME")+":"+viper.GetString("DB_PASSWORD")+"@tcp("+viper.GetString("DB_HOST")+":"+viper.GetString("DB_PORT")+")/"+viper.GetString("DB_DATABASE"))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
