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
	if err := a.db.Table(userTable).Create(&user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}
