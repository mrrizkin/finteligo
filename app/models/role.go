package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Slug      string         `json:"slug" gorm:"unique;not null;index"`
	Name      string         `json:"name"`
}

func (*Role) Seed(db *gorm.DB) {
	data := []Role{
		{Slug: "admin", Name: "Administrator"},
		{Slug: "user", Name: "User"},
	}

	for _, v := range data {
		db.FirstOrCreate(&Role{}, v)
	}
}
