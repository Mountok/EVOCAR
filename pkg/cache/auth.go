package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"todoapp/models"
	"todoapp/pkg/repository"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type AuthRedis struct {
	repos       repository.Authorization
	redisClient *redis.Client
}

func NewAuthCache(repos repository.Authorization, redisClient *redis.Client) *AuthRedis {
	return &AuthRedis{
		repos:       repos,
		redisClient: redisClient,
	}
}

func (r *AuthRedis) CreateUser(user models.User) (userID int, err error) {
	ctx := context.Background()
	key := fmt.Sprintf("users:%s", user.Username)
	TLL := 1 * time.Minute
	userId, err := r.repos.CreateUser(user)
	if err != nil {
		logrus.Errorf("[Cache/Postgres] ошибка при попытке регистрации")
		return 0, err
	}

	// Добавляем пользователя в Redis как Hash
	err = r.redisClient.HSet(ctx, key, "userId", userId, "name", user.Name, "password", user.Password).Err()
	if err != nil {
		logrus.Errorf("Cache] ошибка при добавлении пользователя")
		return 0, err
	}

	// Устанавливаем TTL
	err = r.redisClient.Expire(ctx, key, TLL).Err()
	if err != nil {
		logrus.Errorf("[Cache] ошибка при установке TTL")
		return 0, err
	}

	// Проверяем, что данные записались
	_, err = r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("Ошибка при получении пользователя:", err)
		return
	}
	return userId, nil
}

func (r *AuthRedis) GetUser(username, password string) (models.User, error) {
	ctx := context.Background()
	key := fmt.Sprintf("users:%s", username)
	var user models.User

	// Проверяем TTL ключа
	ttl, err := r.redisClient.TTL(ctx, key).Result()
	if err != nil {
		logrus.Errorf("[Cache] ошибка при получении TTL для username=%s: %s", username, err.Error())
		return user, err
	}

	// Если TTL == -2, значит ключ не существует (истёк)
	if ttl == -2*time.Second {
		logrus.Infof("[Cache] Данные пользователя username=%s отсутствуют в кэше, обращаемся к БД", username)
		return r.repos.GetUser(username, password)
	}

	// Получаем пользователя из Redis
	userData, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		logrus.Errorf("[Cache] ошибка при получении пользователя username=%s ошибка: \n %s", username, err.Error())
		return models.User{}, err
	}

	// Если данных в кэше нет, обращаемся к БД
	if len(userData) == 0 {
		logrus.Infof("[Cache] Пользователь username=%s не найден в кэше, обращаемся к БД", username)
		return r.repos.GetUser(username, password)
	}

	// Парсим данные в структуру models.User
	userId, _ := strconv.Atoi(userData["userId"])
	user = models.User{
		Id:       userId,
		Name:     userData["name"],
		Username: username,
		Password: userData["password"],
	}

	logrus.Infof("[Cache] Пользователь username=%s найден в Redis", username)

	return user, nil
}
