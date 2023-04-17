package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nigelmes/todo/config"
)

const (
	userTable      = "users"
	todoListTable  = "todo_lists"
	usersListTable = "users_list"
	todoItemTable  = "todo_item"
	listItemTable  = "lists_item "
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Dbname,
		cfg.Database.Password,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
