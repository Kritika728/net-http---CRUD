package db

import (
	"database/sql"
	"fmt"
	"net-http---CRUD/pkg/movieactivity"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", movieactivity.DB_USER, movieactivity.DB_PASSWORD, movieactivity.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	movieactivity.CheckErr(err)

	return db
}
