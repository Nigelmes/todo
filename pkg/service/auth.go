package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/repository"
	"time"
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

	return token.SignedString([]byte("sgr^&&fdr(334#F"))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
