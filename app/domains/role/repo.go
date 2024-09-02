package role

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/system/types"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(role *models.Role, permissionIds []uint) error {
	rolePermissions := make([]models.RolePermission, 0)
	for _, permissionId := range permissionIds {
		rolePermissions = append(rolePermissions, models.RolePermission{
			RoleID:       role.ID,
			PermissionID: permissionId,
		})
	}

	err := r.db.Create(&rolePermissions).Error
	if err != nil {
		return err
	}

	return r.db.Create(role).Error
}

func (r *Repo) FindAll(
	pagination types.Pagination,
) ([]models.Role, error) {
	roles := make([]models.Role, 0)
	err := r.db.
		Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&roles).Error
	return roles, err
}

func (r *Repo) FindAllCount() (int64, error) {
	var count int64 = 0
	err := r.db.Model(&models.Role{}).Count(&count).Error
	return count, err
}

func (r *Repo) FindByID(id uint) (*models.Role, error) {
	role := new(models.Role)
	err := r.db.
		Preload("RolePermissions").
		First(role, id).
		Error
	return role, err
}

func (r *Repo) Update(role *models.Role, permissionIds []uint) error {
	err := r.db.
		Unscoped().
		Where("role_id = ?", role.ID).
		Delete(&models.RolePermission{}).
		Error
	if err != nil {
		return err
	}

	rolePermissions := make([]models.RolePermission, 0)
	for _, permissionId := range permissionIds {
		rolePermissions = append(rolePermissions, models.RolePermission{
			RoleID:       role.ID,
			PermissionID: permissionId,
		})
	}

	if len(rolePermissions) == 0 {
		return r.db.Save(role).Error
	}

	err = r.db.Create(&rolePermissions).Error
	if err != nil {
		return err
	}

	return r.db.Save(role).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}
