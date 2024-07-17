package redis

import (
	"user_service/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Load().Server.REDIS_PORT,
		Password: "",
		DB:       0,
	})

	return rdb
}