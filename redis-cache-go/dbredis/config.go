package dbredis

import (
	"redis-cache-go/env"
	"strconv"
)

var RedisConfig = loadConfig()

type config struct {
	Host     string
	Password string
	DB       int
}

func loadConfig() *config {
	c := new(config)

	c.Host = env.GetEnv("REDIS_HOST")
	c.Password = env.GetEnv("REDIS_PASSWORD")
	redisDb := env.GetEnv("REDIS_DB")

	c.DB, _ = strconv.Atoi(redisDb)

	return c
}
