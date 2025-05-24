package repository

import (
	"fmt"
	"todoapp/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)


type AuthPostgres struct {
	db *sqlx.DB
}


func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user models.User) (userID int, err error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id;",usersTable)
	row := r.db.QueryRow(query,user.Name,user.Username, user.Password)
	if err = row.Scan(&id); err != nil {
		logrus.Errorf("[Repository] Ошибка при попытке добавить пользователя в базу: %s", err.Error())
		return 0, err
	} 
	return id, err 
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user,query,username,password)
	if err != nil {
		logrus.Errorf("[Postgres] Ошибка при получении пользователя из БД: %s \n",err.Error())
		return user, err
	}
	logrus.Infof("[Postgres] Пользователь username=%s получен из БД \n",username)
	return user, nil
}