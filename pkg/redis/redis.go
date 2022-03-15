package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/JIeeiroSst/store/model"
	"github.com/go-redis/redis/v8"
)

type RedisDB interface {
	Set(context context.Context, key string, value string) error
	Get(ctx context.Context, key string) (interface{}, error)
	CreateAuth(ctx context.Context, userid string, td *model.TokenDetails) error
	FetchAuth(ctx context.Context, userId string) (string, error)
	FetchToken(ctx context.Context, key string) (string, error)
	DeleteAuth(ctx context.Context, givenUuid string) (int64, error)
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

	if err := r.client.Set(ctx, td.AccessUuid, userid, at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := r.client.Set(ctx, td.RefreshUuid, userid, rt.Sub(now)).Err(); err != nil {
		return err
	}

	accTokenUserId := fmt.Sprintf("%s-%s", at.GoString(), userid)
	rfTokenUserId := fmt.Sprintf("%s-%s", rt.GoString(), userid)

	if err := r.client.Set(ctx, accTokenUserId, td.AccessToken, at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := r.client.Set(ctx, rfTokenUserId, td.RefreshToken, rt.Sub(now)).Err(); err != nil {
		return err
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

func (r *redisDB) FetchToken(ctx context.Context, key string) (string, error) {
	token, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *redisDB) DeleteAuth(ctx context.Context, givenUuid string) (int64, error) {
	deleted, err := r.client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
