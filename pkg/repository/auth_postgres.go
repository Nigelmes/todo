package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nigelmes/todo"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user todo.User) (int, error) {
	if err := a.db.Create(user).Table(userTable).Error; err != nil {
		return -1, err
	}
	return 777, nil
}
