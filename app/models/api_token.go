package models

import (
	"time"

	"gorm.io/gorm"
)

type ApiToken struct {
	ID         uint           `json:"id"          gorm:"primary_key"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"  gorm:"index"`
	Key        string         `json:"key"         gorm:"unique,index"`
	Group      string         `json:"group"`
	Token      string         `json:"token"`
	UserId     uint           `json:"user_id"`
	ExpiryDate *time.Time     `json:"expiry_date"`
	Expired    bool           `json:"expired"`
	Enabled    bool           `json:"enabled"`
	User       User           `json:"user"        gorm:"foreignKey:UserId;references:ID"`
}
