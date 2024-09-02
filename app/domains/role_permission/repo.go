package role_permission

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(role_permission *models.RolePermission) error {
	return r.db.Create(role_permission).Error
}

func (r *Repo) FindAll() ([]models.RolePermission, error) {
	role_permissions := make([]models.RolePermission, 0)
	err := r.db.Find(&role_permissions).Error
	return role_permissions, err
}

func (r *Repo) FindByID(id uint) (*models.RolePermission, error) {
	role_permission := new(models.RolePermission)
	err := r.db.First(role_permission, id).Error
	return role_permission, err
}

func (r *Repo) FindByRoleID(role_id uint) ([]models.RolePermission, error) {
	role_permissions := make([]models.RolePermission, 0)
	err := r.db.Where("role_id = ?", role_id).Find(&role_permissions).Error
	return role_permissions, err
}

func (r *Repo) Update(role_permission *models.RolePermission) error {
	return r.db.Save(role_permission).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.RolePermission{}, id).Error
}
