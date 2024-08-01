package redis

import "github.com/go-redis/redis/v8"

func NewRedis() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	return rds
}
