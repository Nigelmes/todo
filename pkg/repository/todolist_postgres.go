package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nigelmes/todo"
)

type TodoListPostgres struct {
	db *gorm.DB
}

func NewTodoListPostgres(db *gorm.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Table(todoListTable).Create(&list).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Table(usersListTable).Create(&todo.UserList{
		UserId: userId,
		ListId: list.Id,
	}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return list.Id, tx.Commit().Error

}

func (t *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	err := t.db.Table(todoListTable).Joins("INNER JOIN user_lists ON todo_lists.id = user_lists.list_id").
		Where("user_lists.user_id = ?", userId).
		Find(&lists).Error

	return lists, err
}
