package cache

import (
	"context"
	"fmt"
	"log"
	"time"
	"todoapp/pkg/repository"

	"github.com/go-redis/redis/v8"
)

type ConnCache struct {
	repos *repository.Repository
	redisClient *redis.Client
}

func NewConnCache(repos *repository.Repository, redisClient *redis.Client) *ConnCache {
	return &ConnCache{
		repos: repos,
		redisClient: redisClient,
	}
}

func (c *ConnCache) Conn() bool {
	res, err := c.redisClient.Set(context.Background(),"connection","true",time.Minute * 1).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(res)
	return true
}