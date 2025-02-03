package cache

import (
	"todoapp/pkg/repository"

	"github.com/go-redis/redis/v8"
)


type Authorization interface {
	
}

type Conn interface {
	Conn() bool
}


type Cache struct {
	Authorization
	Conn
}

func NewCache(repos *repository.Repository, redisClient *redis.Client) *Cache {
	return &Cache{
		Conn: NewConnCache(repos,redisClient),
	}
}