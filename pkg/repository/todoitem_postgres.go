package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nigelmes/todo"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewTodoItemPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}
	if err := tx.Table(todoItemTable).Create(&item).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Table(listItemTable).Create(&todo.ListsItem{
		ListId: listId,
		ItemId: item.Id,
	}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return item.Id, tx.Commit().Error
}

func (t *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	err := t.db.Table(todoItemTable).Joins("inner join lists_item on lists_item.item_id = todo_items.id").Joins(
		"inner join user_lists on user_lists.list_id = lists_item.list_id").Where("lists_item.list_id = ?"+
		" AND user_lists.user_id = ?", listId, userId).Find(&items).Error
	return items, err
}

func (t *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	err := t.db.Table(todoItemTable).Joins("inner join lists_item on lists_item.item_id = todo_items.id").Joins(
		"inner join user_lists on user_lists.list_id = lists_item.list_id").Where("lists_item.item_id = ?"+
		" AND user_lists.user_id = ?", itemId, userId).First(&item).Error
	return item, err
}

func (t *TodoItemPostgres) Delete(userId, itemId int) error {
	var todoitem todo.TodoItem
	if err := t.db.Where("id = ?", itemId).First(&todoitem).Error; err != nil {
		return fmt.Errorf("failed to delete list with id %d: %w", itemId, err)
	}
	item, err := t.GetById(userId, itemId)
	if err != nil {
		return err
	}
	tx := t.db.Begin()
	if err = tx.Table(todoItemTable).Delete(&item).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Table(listItemTable).Where("item_id = ?", itemId).Delete(&todo.ListsItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (t *TodoItemPostgres) Update(userId, itemId int, input todo.TodoItem) error {
	item, err := t.GetById(userId, itemId)
	if err != nil {
		return err
	}
	err = t.db.Model(&item).Update(input).Error
	return err
}
