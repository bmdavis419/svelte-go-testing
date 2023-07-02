package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() *redis.Client {

	PORT := os.Getenv("REDIS_PORT")
	if PORT == "" {
		PORT = "6379" // Default port if REDIS_PORT environment variable is not set
	}


	rdb := redis.NewClient(&redis.Options{
		// TODO: switch the port to an env variable ✅
		Addr:     "localhost" + PORT,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}