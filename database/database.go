package database

import (
	"database/sql"
	"time"

	"github.com/indraprasetya154/golang-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
