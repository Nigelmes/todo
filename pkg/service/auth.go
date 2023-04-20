package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/repository"
	"time"
)

const (
	signigkey = "sgr^&&fdr(334#Fdfse5644ddd"
)

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	claims := &tokenClaims{
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(12 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		},
		UserId: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signigkey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signigkey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
