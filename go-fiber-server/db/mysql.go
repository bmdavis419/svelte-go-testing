package db

import (
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/config"
	"github.com/jmoiron/sqlx"
)

func CreateMySqlConnection(env config.EnvVars) *sqlx.DB {

	db := sqlx.MustConnect("mysql", env.DSN)

	err := db.Ping()
	if err != nil {
		panic(err)
	} else {
		println("DB CONNECTED")
	}

	return db
}
