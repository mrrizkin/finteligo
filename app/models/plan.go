package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       uint           `json:"price"`
}
