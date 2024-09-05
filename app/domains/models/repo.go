package models

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/system/types"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(langChainLLM *models.LangChainLLM) error {
	return r.db.Create(langChainLLM).Error
}

func (r *Repo) FindAll(
	pagination types.Pagination,
	filter types.Filter,
) ([]models.LangChainLLM, error) {
	langChainLLMs := make([]models.LangChainLLM, 0)
	err := r.db.
		Where(filter.Where, filter.WhereArgs...).
		Offset((pagination.Page - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&langChainLLMs).Error
	if err != nil {
		return nil, err
	}

	return langChainLLMs, err
}

func (r *Repo) FindAllCount(
	filter types.Filter,
) (int64, error) {
	var count int64 = 0
	err := r.db.Model(&models.LangChainLLM{}).
		Where(filter.Where, filter.WhereArgs...).
		Count(&count).
		Error
	return count, err
}

func (r *Repo) FindByID(id uint) (*models.LangChainLLM, error) {
	langChainLLM := new(models.LangChainLLM)
	err := r.db.First(langChainLLM, id).Error
	return langChainLLM, err
}

func (r *Repo) FindByToken(token string) (*models.LangChainLLM, error) {
	langChainLLM := new(models.LangChainLLM)
	err := r.db.Where("token = ?", token).First(langChainLLM).Error
	return langChainLLM, err
}

func (r *Repo) Update(langChainLLM *models.LangChainLLM) error {
	return r.db.Save(langChainLLM).Error
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.LangChainLLM{}, id).Error
}
