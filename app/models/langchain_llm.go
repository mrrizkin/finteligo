package models

import (
	"time"

	"gorm.io/gorm"
)

type LangChainLLM struct {
	ID        uint           `json:"id"         gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UserID    uint           `json:"user_id"`
	Token     string         `json:"token"      gorm:"unique;not null;index"`
	Model     string         `json:"model"                                             validate:"required"`
	Provider  string         `json:"provider"                                          validate:"required"`
	URL       string         `json:"url"`
	APIKey    string         `json:"api_key"`
	Status    string         `json:"status"`
	Enabled   bool           `json:"enabled"                                           validate:"required"`
	Error     string         `json:"error"`
	User      User           `json:"user"       gorm:"foreignKey:UserID;references:ID"`
}
