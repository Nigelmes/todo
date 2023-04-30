package service

import (
	"errors"
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) Delete(userId, listId int) error {
	return t.repo.Delete(userId, listId)
}

func (t *TodoListService) Update(userId, listId int, input todo.TodoList) error {
	if err := validateupdatelist(input); err != nil {
		return err
	}
	return t.repo.Update(userId, listId, input)
}

func validateupdatelist(input todo.TodoList) error {
	if input.Title == "" && input.Description == "" {
		return errors.New("validation error")
	}
	return nil
}
