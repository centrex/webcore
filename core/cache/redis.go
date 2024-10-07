package cache

import (
	"strconv"

	"github.com/centrex/webcore/core/env"
	"github.com/centrex/webcore/core/utils"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	// Get Redis URL.
	redisConn, err := utils.ConnectionURLBuilder("redis")
	dbNumber, _ := strconv.Atoi(env.GetEnv("REDIS_DB", "0"))

	if err != nil {
		panic(err)
	}

	// Set up Redis client.
	return redis.NewClient(&redis.Options{
		Addr:     redisConn,
		Password: env.GetEnv("REDIS_PASSWORD", ""),
		DB:       dbNumber,
	})
}
