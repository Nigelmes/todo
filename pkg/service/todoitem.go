package service

import (
	"errors"
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/repository"
)

type TodoItemService struct {
	repo_list repository.TodoList
	repo_item repository.TodoItem
}

func NewTodoItemService(repo_list repository.TodoList, repo_item repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo_list: repo_list, repo_item: repo_item}
}

func (t *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	if _, err := t.repo_list.GetById(userId, listId); err != nil {
		return 0, err
	}
	return t.repo_item.Create(listId, item)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return t.repo_item.GetAll(userId, listId)
}

func (t *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return t.repo_item.GetById(userId, itemId)
}

func (t *TodoItemService) Delete(userId, itemId int) error {
	return t.repo_item.Delete(userId, itemId)
}

func (t *TodoItemService) Update(userId, itemId int, input todo.TodoItem) error {
	if err := validateupdateitem(input); err != nil {
		return err
	}
	return t.repo_item.Update(userId, itemId, input)
}

func validateupdateitem(input todo.TodoItem) error {
	if input.Title == "" && input.Description == "" {
		return errors.New("validation error")
	}
	return nil
}
