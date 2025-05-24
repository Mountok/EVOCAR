package service

import (
	"todoapp/models"
	"todoapp/pkg/cache"
)

type OrderService struct {
	cache cache.Order
}

func NewOrderService(cache cache.Order) *OrderService {
	return &OrderService{cache: cache}
}

func (s *OrderService) GetOrders() ([]models.Order, error) {
	return s.cache.GetOrders()
}
func (s *OrderService) GetActiveOrders() ([]models.Order, error) {
	return s.cache.GetActiveOrders()
}

func (s *OrderService) 	CompleteOrder(id string) error {
	return s.cache.CompleteOrder(id)
}

func (s *OrderService) CreateOrder(order models.Order) (int, error) {
	return s.cache.CreateOrder(order)
}

func (s *OrderService) AcceptOrder(id string, phoneNumber string) error {
	return s.cache.AcceptOrder(id, phoneNumber)
}
func (s *OrderService) CancleOrder(id string) error {
	return s.cache.CancleOrder(id)
}


func (s *OrderService) GetOrdersByPhoneNumber(phoneNumber string) ([]models.Order, error) {
    return s.cache.GetOrdersByPhoneNumber(phoneNumber)
}

func (s *OrderService) GetExecutorsHistory(phoneNumber string) ([]models.Order, error) {
	return s.cache.GetExecutorsHistory(phoneNumber)
}