package models

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Code      string         `json:"code" gorm:"unique;not null;index"`
	Name      string         `json:"name"`
}

func (*Status) Seed(db *gorm.DB) {
	statuses := []Status{
		{Code: "active", Name: "Active"},
		{Code: "inactive", Name: "Inactive"},
		{Code: "pending", Name: "Pending"},
		{Code: "blocked", Name: "Blocked"},
		{Code: "deleted", Name: "Deleted"},
		{Code: "completed", Name: "Completed"},
		{Code: "canceled", Name: "Canceled"},
		{Code: "expired", Name: "Expired"},
		{Code: "suspended", Name: "Suspended"},
		{Code: "rejected", Name: "Rejected"},
		{Code: "approved", Name: "Approved"},
		{Code: "waiting", Name: "Waiting"},
		{Code: "processing", Name: "Processing"},
		{Code: "refunded", Name: "Refunded"},
		{Code: "failed", Name: "Failed"},
		{Code: "success", Name: "Success"},
		{Code: "error", Name: "Error"},
		{Code: "warning", Name: "Warning"},
		{Code: "info", Name: "Info"},
		{Code: "danger", Name: "Danger"},
	}

	for _, status := range statuses {
		db.FirstOrCreate(&Status{}, status)
	}
}
