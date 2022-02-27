package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisDB interface {
	Set(context context.Context, key string, value string) error
	Get(ctx context.Context, key string) (interface{}, error)
}

type redisDB struct {
	client *redis.Client
}

func NewDatabase(address string) RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return &redisDB{
		client: client,
	}
}

func (r *redisDB) Set(ctx context.Context, key string, value string) error {
	if err := r.client.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisDB) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
