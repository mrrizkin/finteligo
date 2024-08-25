package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	PlanID    uint           `json:"plan_id"`
	UserID    uint           `json:"user_id"`
	StartDate string         `json:"start_date" gorm:"index"`
	EndDate   string         `json:"end_date" gorm:"index"`
	StatusID  uint           `json:"status_id"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Plan      Plan           `json:"plan" gorm:"foreignKey:PlanID;references:ID"`
	Status    Status         `json:"status" gorm:"foreignKey:StatusID;references:ID"`
}
