package service

import (
	"todoapp/models"
	"todoapp/pkg/cache"
)

type Authorization interface {
	CreateUser(user models.User) (userId int, err error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Conn interface {
	Conn() bool
}

type Order interface {
	GetOrders() ([]models.Order, error)
	CreateOrder(order models.Order) (int, error)
	AcceptOrder(id string, phoneNumber string) error
	CompleteOrder(id string) error
	GetOrdersByPhoneNumber(phoneNumber string) ([]models.Order, error)
	GetActiveOrders() ([]models.Order, error)
	CancleOrder(id string) error
	GetExecutorsHistory(phoneNumber string) ([]models.Order, error)
	CheckOrderStatus(orderId string) (map[string]interface{}, error)
	GetOrderExecutorById(id string) (models.ExecutorHistory, error)
}

type Service struct {
	Authorization
	Conn
	Order
}

func NewService(cache *cache.Cache) *Service {
	return &Service{
		Conn:          NewConnService(cache),
		Authorization: NewAuthService(cache),
		Order:         NewOrderService(cache),
	}
}
