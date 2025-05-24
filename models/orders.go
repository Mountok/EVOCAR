package models

import "time"

type Order struct {
	Id             int       `json:"id" db:"id"`
	Latitude       float64   `json:"latitude" db:"latitude" binding:"required"`
	Longitude      float64   `json:"longitude" db:"longitude" binding:"required"`
	Location       *string   `json:"location" db:"location" binding:"required"`
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
