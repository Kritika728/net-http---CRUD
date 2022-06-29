package db

import (
	"database/sql"
	"fmt"
	"nethttpcrud/lib"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", lib.DB_USER, lib.DB_PASSWORD, lib.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	lib.CheckErr(err)

	return db
}
