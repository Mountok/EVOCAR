package models

import "time"

type Order struct {
	Id             int       `json:"id" db:"id"`
	FromLatitude   float64   `json:"from_latitude" db:"from_latitude" binding:"required"`
	FromLongitude  float64   `json:"from_longitude" db:"from_longitude" binding:"required"`
	FromLocation   *string   `json:"from_location" db:"from_location" binding:"required"`
	ToLatitude     float64   `json:"to_latitude" db:"to_latitude"`
	ToLongitude    float64   `json:"to_longitude" db:"to_longitude"`
	ToLocation     *string   `json:"to_location" db:"to_location"`
	TypeOfOrder    string    `json:"typeOfOrder" db:"typeoforder" binding:"required"`
	TypeOfAuto     string    `json:"typeOfAuto" db:"typeofauto" binding:"required"`
	NumberOfClient string    `json:"numberOfClient" db:"numberofclient" binding:"required"`
	Status         string    `json:"status" db:"status"  binding:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type ExecutorHistory struct {
	Id             int       `json:"id" db:"id"`
	OrderId        int       `json:"orderId" db:"order_id"`
	ExecutorNumber string    `json:"executorNumber" db:"executor_number"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
