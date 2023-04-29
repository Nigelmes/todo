package repository

import (
	"fmt"
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

func (t *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	err := t.db.Table(todoListTable).Joins("INNER JOIN user_lists ON todo_lists.id = user_lists.list_id").
		Where("user_lists.user_id = ? AND user_lists.list_id = ?", userId, listId).
		Find(&list).Error

	return list, err
}

func (t *TodoListPostgres) Delete(userId, listId int) error {
	var todoList todo.TodoList
	if err := t.db.Where("id = ?", listId).First(&todoList).Error; err != nil {
		return fmt.Errorf("failed to find list with id %d: %w", listId, err)
	}
	tx := t.db.Begin()
	if err := tx.Where("user_id = ? AND list_id = ?", userId, listId).Delete(&todo.UserList{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id = ?", listId).Delete(&todo.TodoList{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
