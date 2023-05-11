package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://localhost:5432/go_svelte_todos?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func SeedDatabase(db *sqlx.DB) {
	tx := db.MustBegin()
	insertStatement := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)"
	tx.MustExec(insertStatement, "John", "Doe", "john.doe@gmail.com", "1212323412")
	tx.MustExec(insertStatement, "Ben", "Davis", "test@gmail.com", "2112390123")
	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}
