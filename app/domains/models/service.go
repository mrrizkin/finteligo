package models

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/third_party/langchain"
)

func NewService(repo *Repo, lc *langchain.LangChain) *Service {
	return &Service{
		repo:      repo,
		langchain: lc,
	}
}

func (s *Service) Create(langChainLLM *models.LangChainLLM) (*models.LangChainLLM, error) {
	err := s.langchain.AddLLM(langchain.AddLLMParams{
		Token:    langChainLLM.Token,
		Model:    langChainLLM.Model,
		Provider: langChainLLM.Provider,
		URL:      langChainLLM.URL,
		APIKey:   langChainLLM.APIKey,
	})
	if err != nil {
		return nil, err
	}

	lcLLM, err := s.repo.FindByToken(langChainLLM.Token)
	if err != nil {
		return nil, err
	}

	return lcLLM, nil
}

func (s *Service) FindAll() ([]models.LangChainLLM, error) {
	langChainLLMs, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return langChainLLMs, nil
}

func (s *Service) FindByID(id uint) (*models.LangChainLLM, error) {
	langChainLLM, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return langChainLLM, nil
}

// TODO: need to implement langchain update
func (s *Service) Update(id uint, langChainLLM *models.LangChainLLM) (*models.LangChainLLM, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(langChainLLM)
	if err != nil {
		return nil, err
	}

	return langChainLLM, nil
}

func (s *Service) Delete(id uint) error {
	lcLLM, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.langchain.RemoveLLM(lcLLM.Token)
}
