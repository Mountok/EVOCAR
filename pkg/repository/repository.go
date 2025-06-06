package repository

import (
	"todoapp/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (userID int, err error)
	GetUser(username, password string) (models.User, error)
}

type Conn interface {
}

type Orders interface {
	CreateOrder(order models.Order) (int, error)
	GetOrders() ([]models.Order, error)
	AcceptOrder(id string, phoneNumber string) error
	CompleteOrder(id string) error
	GetOrdersByPhoneNumber(phoneNumber string) ([]models.Order, error)
	GetActiveOrders() ([]models.Order, error)
	CancleOrder(id string) error
	GetExecutorsHistory(phoneNumber string) ([]models.Order, error)
	CheckOrderStatus(orderId string) (string, error)
	GetOrderExecutorById(id int) (models.ExecutorHistory, error)
}

type Repository struct {
	Authorization
	Conn
	Orders
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Orders:        NewOrdersPostgres(db),
	}
}
