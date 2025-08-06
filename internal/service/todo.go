package service

import (
	"api/internal/models"
	"api/internal/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todoItem models.TodoItem) {
	//todoItem.Id = uuid.New()
	s.repo.Create(&todoItem)
}

func (s *TodoService) Get(id uuid.UUID) *models.TodoItem {
	return s.repo.Get(id)
}
