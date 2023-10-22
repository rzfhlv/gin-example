package redis

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	once        sync.Once
	redisError  error
)

type Redis struct {
	client *redis.Client
}

func New() (*Redis, error) {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			redisError = err
		}
	})

	if redisError != nil {
		return nil, redisError
	}

	return &Redis{
		client: redisClient,
	}, nil
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}
