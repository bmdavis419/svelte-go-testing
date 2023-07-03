package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DSN            string
	REDIS_ADDR     string
	REDIS_PASSWORD string
	REDIS_DB       int
}

// this is cringe in its current state but it works
func LoadEnv() EnvVars {
	godotenv.Load()

	dsn := os.Getenv("DSN")
	redis_addr := os.Getenv("REDIS_ADDR")
	redis_pass := os.Getenv("REDIS_PASSWORD")
	redis_db := os.Getenv("REDIS_DB")
	parsed_redis_db, err := strconv.Atoi(redis_db)
	if err != nil {
		panic("cannot parse redis DB number")
	}

	return EnvVars{
		DSN:            dsn,
		REDIS_ADDR:     redis_addr,
		REDIS_PASSWORD: redis_pass,
		REDIS_DB:       parsed_redis_db,
	}
}
