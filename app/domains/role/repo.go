package role

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *Repo) FindAll() ([]models.Role, error) {
	roles := make([]models.Role, 0)
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *Repo) FindByID(id uint) (*models.Role, error) {
	role := new(models.Role)
	err := r.db.First(role, id).Error
	return role, err
}

func (r *Repo) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}
