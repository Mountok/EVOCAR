package cache

import (
	"todoapp/models"
	"todoapp/pkg/repository"

	"github.com/go-redis/redis/v8"
)

type Authorization interface {
	CreateUser(user models.User) (userID int, err error)
	GetUser(username, password string) (models.User, error)
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
	CheckOrderStatus(orderId string) (string, error)
	GetOrderExecutorById(id int) (models.ExecutorHistory, error)
}

type Cache struct {
	Authorization
	Conn
	Order
}

func NewCache(repos *repository.Repository, redisClient *redis.Client) *Cache {
	return &Cache{
		Conn:          NewConnCache(repos, redisClient),
		Authorization: NewAuthCache(repos, redisClient),
		Order:         NewOrderCache(repos, redisClient),
	}
}
