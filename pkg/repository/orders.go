package repository

import (
	"todoapp/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrdersPostgres struct {
	db *sqlx.DB
}

func NewOrdersPostgres(db *sqlx.DB) *OrdersPostgres {
	return &OrdersPostgres{db: db}
}

func (r *OrdersPostgres) CreateOrder(order models.Order) (int, error) {
	query := `INSERT INTO orders (latitude, longitude, location, typeOfOrder, numberOfClient,typeOfAuto, status)
			 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err := r.db.QueryRow(query,
		order.Latitude,
		order.Longitude,
		order.Location,
		order.TypeOfOrder,
		order.NumberOfClient,
		order.TypeOfAuto,
		order.Status).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrdersPostgres) GetOrders() ([]models.Order, error) {
	query := `SELECT * FROM orders;`
	var orders []models.Order
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (r *OrdersPostgres) GetActiveOrders() ([]models.Order, error) {
	query := `SELECT * FROM orders WHERE status='ожидание';`
	var orders []models.Order
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrdersPostgres) AcceptOrder(id string, phoneNumber string) error {
	query := `UPDATE orders SET status = 'принят' WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	logrus.Infoln("Заказ принят", id, phoneNumber)

	query = `INSERT INTO executors_orders_history (order_id, executor_number) VALUES ($1, $2)`
	_, err = r.db.Exec(query, id, phoneNumber)
	if err != nil {
		return err
	}
	logrus.Infoln("Запрос на вставку в историю", query, id, phoneNumber)

	return err
}

func (r *OrdersPostgres) CompleteOrder(id string) error {
	query := `UPDATE orders SET status = 'выполнен' WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	logrus.Infoln("Заказ %s выпол", id)
	return err
}

func (r *OrdersPostgres) CancleOrder(id string) error {
	query := `UPDATE orders SET status = 'отменен' WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *OrdersPostgres) GetOrdersByPhoneNumber(phoneNumber string) ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, latitude, longitude, location, typeOfOrder, numberOfClient, status, created_at
              FROM orders WHERE numberOfClient = $1`

	err := r.db.Select(&orders, query, phoneNumber)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrdersPostgres) GetExecutorsHistory(phoneNumber string) ([]models.Order, error) {
	var history []models.ExecutorHistory
	query := `SELECT id, order_id, executor_number, created_at
			  FROM executors_orders_history WHERE executor_number = $1`
	err := r.db.Select(&history, query, phoneNumber)
	if err != nil {
		return nil, err
	}
	// Получить заказы по id из истории
	var orders []models.Order
	for _, h := range history {
		var order models.Order
		query := `SELECT * FROM orders WHERE id = $1`
		err := r.db.Get(&order, query, h.OrderId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
