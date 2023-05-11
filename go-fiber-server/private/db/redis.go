package db

import (
	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		// TODO: switch the port to an env variable
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
