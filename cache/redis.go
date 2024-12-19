package cache

import (
	"context"
	"gokweb/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPwd, // no password set
		DB:       config.RedisDb,  // use default DB
	})

	Rdb = rdb
	Rctx = ctx
}
