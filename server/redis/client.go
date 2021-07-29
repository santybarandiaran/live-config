package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"live-config/logger"
)

const defaultTopicTemplate = "live-config_%s_%s_%s"

var (
	Redis *redis.Client
	ctx = context.Background()
)

func Init() *redis.Client {
	conf := New()

	client := redis.NewClient(&redis.Options{
		Addr:       *conf.Address,
		//Password:   *conf.Password,
		Password:   "",
		OnConnect:  logConnection,
		DB:         0,
		MaxRetries: 10,
	})

	Redis = client

	return Redis
}

func logConnection(_ context.Context, _ *redis.Conn) error {
	logger.Instance.Info("connection to redis established")
	return nil
}
