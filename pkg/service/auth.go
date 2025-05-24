package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"todoapp/models"
	"todoapp/pkg/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const salt = "dvr39g8ehdfgn54"
const signingKey = "ef32ewrf65rf"
const tokenTTL = 2 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	cache cache.Authorization
}

func NewAuthService(cache cache.Authorization) *AuthService {
	return &AuthService{
		cache: cache,
	}
}

func (s *AuthService) CreateUser(user models.User) (userID int, err error) {
	user.Password = generatePassword(user.Password)
	return s.cache.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.cache.GetUser(username, generatePassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id,
		})
	logrus.Infof("[Service] Токен для пользователя с username=%s сгенерирован", username)
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный ключ подписи")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID,nil
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
