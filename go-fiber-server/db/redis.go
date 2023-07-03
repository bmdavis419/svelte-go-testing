package db

import (
	"github.com/bmdavis419/svelte-go-testing/go-fiber-server/config"
	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection(env config.EnvVars) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_ADDR,
		Password: env.REDIS_PASSWORD,
		DB:       env.REDIS_DB,
	})

	return rdb
}
