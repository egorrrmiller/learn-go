package models

import (
	"github.com/google/uuid"
)

type TodoItem struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string
	Description string
}
