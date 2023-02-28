package cache

import "github.com/redis/go-redis/v9"

func Connection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
