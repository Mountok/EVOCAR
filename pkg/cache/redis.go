package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisClient(cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
		// TLSConfig: &tls.Config{},
	})
	
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
