package repository

import (
	"api/internal/models"

	"github.com/google/uuid"
)

type TodoItemRepository interface {
	Create(todoItem models.TodoItem) uuid.UUID
	Get(id uuid.UUID) models.TodoItem
}
