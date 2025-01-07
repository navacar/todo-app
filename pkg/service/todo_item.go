package service

import (
	todo "rest-api-todo"
	"rest-api-todo/pkg/repository"
)

type TodoItemtService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemtService {
	return &TodoItemtService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemtService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemtService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemtService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemtService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemtService) Update(userId int, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
