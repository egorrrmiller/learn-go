package dto

import "github.com/google/uuid"

type TodoItemRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoItemResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
