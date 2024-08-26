package models

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

func NewRepo(db *database.Database) *Repo {
	return &Repo{db}
}

func (r *Repo) Create(langChainLLM *models.LangChainLLM) error {
	return r.db.Create(langChainLLM).Error
}

func (r *Repo) FindAll() ([]models.LangChainLLM, error) {
	langChainLLMs := make([]models.LangChainLLM, 0)
	err := r.db.Find(&langChainLLMs).Error
	return langChainLLMs, err
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
