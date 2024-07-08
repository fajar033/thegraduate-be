package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type RedisClient struct {
	redis *redis.Client
}

func NewRedisClient() *RedisClient {
	var redisClient *redis.Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       10,
	})

	return &RedisClient{
		redis: redisClient,
	}
}
func (r *RedisClient) GetValue(key string) (string, error) {

	var rdbCommand *redis.StringCmd = r.redis.Get(context.Background(), key)

	if err := rdbCommand.Err(); err != nil {

		return "", err
	}

	data, err := rdbCommand.Result()

	if err != nil {
		return "", err
	}

	return data, nil

}

func (r *RedisClient) SetValue(key string, value string, ttl time.Duration) {

	var rdbComamnd *redis.StatusCmd = r.redis.Set(context.Background(), key, value, ttl)
	if err := rdbComamnd.Err(); err != nil {
		panic(err)
	}

}
