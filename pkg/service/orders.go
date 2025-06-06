package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
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

func (s *OrderService) CompleteOrder(id string) error {
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

func (s *OrderService) GetOrderExecutorById(id string) (models.ExecutorHistory, error) {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return models.ExecutorHistory{}, err
	}
	result, err := s.cache.GetOrderExecutorById(orderId)
	if err != nil {
		return models.ExecutorHistory{}, err
	}
	return result, nil
}

func (s *OrderService) CheckOrderStatus(orderId string) (map[string]interface{}, error) {
	if orderId == "" || orderId == "0" || orderId == " " {
		return nil, errors.New("orderId is zero")
	}

	status, err := s.cache.CheckOrderStatus(orderId)
	if err != nil {
		return nil, err
	}
	if status == "принят" {
		logrus.Println("Статус заказа %s изменился на принят")
		result, err := s.GetOrderExecutorById(orderId)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"orderId":       orderId,
			"status":        status,
			"executor_data": result,
		}, err
	}

	logrus.Printf("Status of order %s: %s \n", orderId, status)
	return map[string]interface{}{
		"status":  status,
		"orderId": orderId,
	}, nil
}
