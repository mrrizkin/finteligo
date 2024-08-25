package user

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repo) FindAll() ([]models.User, error) {
	users := make([]models.User, 0)
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repo) FindByID(id uint) (*models.User, error) {
	user := new(models.User)
	err := r.db.First(user, id).Error
	return user, err
}

func (r *Repo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
