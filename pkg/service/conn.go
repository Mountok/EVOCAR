package service

import "todoapp/pkg/cache"

type ConnService struct {
	cache *cache.Cache
}

func NewConnService(cache *cache.Cache) *ConnService {
	return &ConnService{
		cache: cache,
	}
} 


func (s *ConnService) Conn() bool {
	return s.cache.Conn.Conn()
}