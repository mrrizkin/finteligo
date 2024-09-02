package api_token

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/system/types"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(apiToken *models.ApiToken) error {
	return r.db.Create(apiToken).Error
}

func (r *Repo) FindAll(
	pagination types.Pagination,
) ([]models.ApiToken, error) {
	apiTokens := make([]models.ApiToken, 0)
	err := r.db.
		Preload("User").
		Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&apiTokens).Error
	if err != nil {
		return nil, err
	}

	return apiTokens, err
}

func (r *Repo) FindAllCount() (int64, error) {
	var count int64 = 0
	err := r.db.Model(&models.ApiToken{}).Count(&count).Error
	return count, err
}

func (r *Repo) FindByID(id uint) (*models.ApiToken, error) {
	apiToken := new(models.ApiToken)
	err := r.db.First(apiToken, id).Error
	return apiToken, err
}

func (r *Repo) FindByToken(token string) (*models.ApiToken, error) {
	apiToken := new(models.ApiToken)
	err := r.db.Where("token = ?", token).First(apiToken).Error
	return apiToken, err
}

func (r *Repo) Enable(id uint) (*models.ApiToken, error) {
	apiToken := new(models.ApiToken)
	err := r.db.First(apiToken, id).Error
	if err != nil {
		return nil, err
	}

	apiToken.Enabled = true
	err = r.Update(apiToken)
	return apiToken, err
}

func (r *Repo) Disable(id uint) (*models.ApiToken, error) {
	apiToken := new(models.ApiToken)
	err := r.db.First(apiToken, id).Error
	if err != nil {
		return nil, err
	}

	apiToken.Enabled = false
	err = r.Update(apiToken)
	return apiToken, err
}

func (r *Repo) Update(apiToken *models.ApiToken) error {
	return r.db.Save(apiToken).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.ApiToken{}, id).Error
}
