package auth

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := r.db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
