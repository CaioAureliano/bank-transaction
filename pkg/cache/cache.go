package cache

import (
	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: configuration.Env.REDISURI,
		DB:   0,
	})
}
