package service

import "todoapp/pkg/cache"


type Authorization interface {
	
}

type Conn interface {
	Conn() bool
}


type Service struct {
	Authorization
	Conn
}

func NewService(cache *cache.Cache) *Service {
	return &Service{
		Conn: NewConnService(cache),
	}
}