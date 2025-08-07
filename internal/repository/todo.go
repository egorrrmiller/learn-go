package repository

import (
	"api/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	err := db.AutoMigrate(&models.TodoItem{})

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo models.TodoItem) {
	r.db.Create(&todo)
}

func (r *TodoRepository) Get(id uuid.UUID) models.TodoItem {
	todo := models.TodoItem{Id: id}

	r.db.First(&todo)

	return todo
}
