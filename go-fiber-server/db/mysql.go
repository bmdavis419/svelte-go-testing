package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func CreateMySqlConnection() *sqlx.DB {
	conn_str := os.Getenv("DSN")

	db := sqlx.MustConnect("mysql", conn_str)

	err := db.Ping()
	if err != nil {
		panic(err)
	} else {
		println("DB CONNECTED")
	}

	return db
}
