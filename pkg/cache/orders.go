package cache

import (
	"context"
	"todoapp/models"
	"todoapp/pkg/repository"

	"github.com/go-redis/redis/v8"
)

type OrderCache struct {
	repos repository.Orders
	redisClient *redis.Client
}

func NewOrderCache(repos repository.Orders, redisClient *redis.Client) *OrderCache {
	return &OrderCache{repos: repos, redisClient: redisClient}
}

func (c *OrderCache) GetOrders() ([]models.Order, error) {
	return c.repos.GetOrders()
}
func (c *OrderCache) GetActiveOrders() ([]models.Order, error) {
	return c.repos.GetActiveOrders()
}
func (c *OrderCache) CreateOrder(order models.Order) (int, error) {
	id, err := c.repos.CreateOrder(order)
	if err != nil {
		return 0, err
	}
	c.redisClient.Set(context.Background(), "orders", order, 0)
	return id, nil
}

func (c *OrderCache) AcceptOrder(id string, phoneNumber string) error {
	return c.repos.AcceptOrder(id, phoneNumber)
}

func (c *OrderCache) CompleteOrder(id string) error {
	return c.repos.CompleteOrder(id)
}

func (c *OrderCache) CancleOrder(id string) error {
	return c.repos.CancleOrder(id)
}



func (c *OrderCache) GetOrdersByPhoneNumber(phoneNumber string)  ([]models.Order, error) {
	return c.repos.GetOrdersByPhoneNumber(phoneNumber)
}

func (c *OrderCache) GetExecutorsHistory(phoneNumber string) ([]models.Order, error) {
	return c.repos.GetExecutorsHistory(phoneNumber)
}
