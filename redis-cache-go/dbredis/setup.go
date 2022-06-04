package dbredis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	Redis = new(RedisDB)
	ctx   = context.TODO()
)

type RedisDB struct {
	Client *redis.Client
	Cache  *cache.Cache
}

func RedisStart() (*RedisDB, error) {
	rs := new(RedisDB)
	var err error

	rs.Client = redis.NewClient(&redis.Options{
		Addr:     RedisConfig.Host,
		Network:  "tcp",
		Password: RedisConfig.Password, // no password set
		DB:       RedisConfig.DB,       // use default DB
	})

	rs.Cache = cache.New(&cache.Options{
		Redis: rs.Client,
	})

	if err = rs.Client.Set(ctx, "key", "value", time.Minute).Err(); err != nil {
		defer rs.Client.Close()
		panic(err)
	}

	Redis = rs

	fmt.Println("Redis starting...")

	return rs, nil

}

func (r *RedisDB) CheckExistKey(key string) bool {
	return r.Cache.Exists(ctx, key)
}

func (r *RedisDB) SetCache(key string, value interface{}, expired time.Duration) {
	if err := r.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   expired,
	}); err != nil {
		panic(err)
	}

	return
}

func (r *RedisDB) GetCache(key string, value interface{}) error {
	return r.Cache.Get(ctx, key, &value)
}
