package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Slug      string         `json:"slug" gorm:"unique;not null;index"`
	Name      string         `json:"name"`
}

func (*Permission) Seed(db *gorm.DB) {
	data := []Permission{
		{Slug: "create_database", Name: "Create Database"},
		{Slug: "read_database", Name: "Read Database"},
		{Slug: "update_database", Name: "Update Database"},
		{Slug: "delete_database", Name: "Delete Database"},

		{Slug: "create_job", Name: "Create Job"},
		{Slug: "read_job", Name: "Read Job"},
		{Slug: "update_job", Name: "Update Job"},
		{Slug: "delete_job", Name: "Delete Job"},

		{Slug: "create_permission", Name: "Create Permission"},
		{Slug: "read_permission", Name: "Read Permission"},
		{Slug: "update_permission", Name: "Update Permission"},
		{Slug: "delete_permission", Name: "Delete Permission"},

		{Slug: "create_role_permission", Name: "Create Role Permission"},
		{Slug: "read_role_permission", Name: "Read Role Permission"},
		{Slug: "update_role_permission", Name: "Update Role Permission"},
		{Slug: "delete_role_permission", Name: "Delete Role Permission"},

		{Slug: "create_role", Name: "Create Role"},
		{Slug: "read_role", Name: "Read Role"},
		{Slug: "update_role", Name: "Update Role"},
		{Slug: "delete_role", Name: "Delete Role"},

		{Slug: "create_user", Name: "Create User"},
		{Slug: "read_user", Name: "Read User"},
		{Slug: "update_user", Name: "Update User"},
		{Slug: "delete_user", Name: "Delete User"},
	}

	for _, v := range data {
		db.FirstOrCreate(&Permission{}, v)
	}
}
