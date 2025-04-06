package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: rdb,
	}
}

func (r *RedisStorage) Increment(key string, expiration time.Duration) (int, error) {
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if val == 1 {
		r.client.Expire(ctx, key, expiration)
	}

	return int(val), nil
}

func (r *RedisStorage) Get(key string) (int, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	var result int
	fmt.Sscanf(val, "%d", &result)
	return result, nil
}

func (r *RedisStorage) Block(key string, duration time.Duration) error {
	return r.client.Set(ctx, key+":block", 1, duration).Err()
}

func (r *RedisStorage) IsBlocked(key string) (bool, error) {
	val, err := r.client.Exists(ctx, key+":block").Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

func (r *RedisStorage) Reset(key string) error {
	return r.client.Del(ctx, key).Err()
}
