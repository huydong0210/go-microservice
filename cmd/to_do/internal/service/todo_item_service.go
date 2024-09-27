package service

import (
	"go-microservices/cmd/to_do/internal/repository"
	"go-microservices/cmd/to_do/pkg/model"
)

type TodoItemServiceInterface interface {
	CreateTodoItem(item *model.TodoItem) error
	DeleteTodoItem(itemId uint, userId uint) error
	UpdateTodoItem(itemId uint, userId uint, item *model.TodoItem) error
	FindTodoItemById(itemId uint, userId uint) (*model.TodoItem, error)
}
type TodoItemService struct {
	repo *repository.TodoItemRepository
}

func NewTodoItemService(todoItemRepo *repository.TodoItemRepository) *TodoItemService {
	return &TodoItemService{repo: todoItemRepo}
}
func (s TodoItemService) CreateTodoItem(item *model.TodoItem) error {
	return s.repo.CreateToDoItem(item)
}
func (s TodoItemService) DeleteTodoItem(itemId uint, userId uint) error {
	return s.repo.DeleteTodoItem(itemId, userId)
}
func (s TodoItemService) UpdateTodoItem(itemId uint, userId uint, item *model.TodoItem) error {
	return s.repo.UpdateTodoItem(itemId, userId, item)
}
func (s TodoItemService) FindTodoItemById(itemId uint, userId uint) (*model.TodoItem, error) {
	return s.repo.FindTodoItemById(itemId, userId)
}
