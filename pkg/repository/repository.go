package repository

import "github.com/jmoiron/sqlx"


type Authorization interface {
	
}

type Conn interface {

}


type Repository struct {
	Authorization
	Conn
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}