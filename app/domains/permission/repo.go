package permission

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

func (r *Repo) FindAll() ([]models.Permission, error) {
	permissions := make([]models.Permission, 0)
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *Repo) FindByID(id uint) (*models.Permission, error) {
	permission := new(models.Permission)
	err := r.db.First(permission, id).Error
	return permission, err
}

func (r *Repo) Update(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Permission{}, id).Error
}
