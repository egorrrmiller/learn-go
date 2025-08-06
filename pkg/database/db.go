package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "1", "postgres", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
