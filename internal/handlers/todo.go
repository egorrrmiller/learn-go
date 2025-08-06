package handlers

import (
	"api/internal/models"
	"api/internal/models/dto"
	"api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var request dto.TodoItemRequest

	if err := c.BindJSON(&request); err != nil {
		logrus.Error("Invalid body")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	model := models.TodoItem{
		Title:       request.Title,
		Description: request.Description,
	}

	h.service.Create(model)

	c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) Get(c *gin.Context) {

	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		logrus.Error("Invalid body")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response := h.service.Get(todoId)

	result := dto.TodoItemResponse{
		Id:          response.Id,
		Title:       response.Title,
		Description: response.Description,
	}

	c.JSON(http.StatusOK, result)
}
