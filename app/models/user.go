package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"         gorm:"primary_key"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  *string        `json:"username"   gorm:"unique;not null;index"`
	Password  *string        `json:"password"`
	Name      *string        `json:"name"`
	Email     *string        `json:"email"`
	RoleID    *uint          `json:"role_id"`
	Role      Role           `json:"role"       gorm:"foreignKey:RoleID;references:ID"`
}

func (*User) Seed(db *gorm.DB) {
	var adminRole Role
	db.Where("slug = ?", "admin").First(&adminRole)

	username := "super_mimin"
	password := "super_mimin"
	email := "m-super-admin@mail.com"
	name := "Administrator"

	user := User{
		Username: &username,
		Password: &password,
		Email:    &email,
		Name:     &name,
		RoleID:   &adminRole.ID,
	}

	db.FirstOrCreate(&User{}, user)
}
