package main

import "github.com/bmdavis419/svelte-go-testing/go-fiber-server/db"

func main() {
	conn := db.CreateConnection()

	db.SeedDatabase(conn)
}
