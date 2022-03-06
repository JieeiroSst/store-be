package redis

import (
	"context"
	"time"

	"github.com/JIeeiroSst/store/model"
	"github.com/go-redis/redis/v8"
)

type RedisDB interface {
	Set(context context.Context, key string, value string) error
	Get(ctx context.Context, key string) (interface{}, error)
	CreateAuth(ctx context.Context, userid string, td *model.TokenDetails) error
	FetchAuth(ctx context.Context, userId string) (string, error)
	DeleteAuth(ctx context.Context, givenUuid string) (int64, error)
	UpdateAuth(ctx context.Context, key string, value string) error
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

func (r *redisDB) CreateAuth(ctx context.Context, userid string, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := r.client.Set(ctx, td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := r.client.Set(ctx, td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (r *redisDB) FetchAuth(ctx context.Context, userId string) (string, error) {
	userid, err := r.client.Get(ctx, userId).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (r *redisDB) DeleteAuth(ctx context.Context, givenUuid string) (int64, error) {
	deleted, err := r.client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (r *redisDB) UpdateAuth(ctx context.Context, key string, value string) error {
		

	return nil
}
