package models

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
	"github.com/mrrizkin/finteligo/third_party/langchain"
	lcTypes "github.com/mrrizkin/finteligo/third_party/langchain/types"
)

func NewService(repo *Repo, lc *langchain.LangChain) *Service {
	return &Service{
		repo:      repo,
		langchain: lc,
	}
}

func (s *Service) Create(langChainLLM *models.LangChainLLM) (*models.LangChainLLM, error) {
	err := s.langchain.AddLLM(lcTypes.AddLLMParams{
		UserID:   langChainLLM.UserID,
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

func (s *Service) FindAll(
	pagination types.Pagination,
) (*PaginatedModels, error) {
	langChainLLMs, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	langChainLLMsCount, err := s.repo.FindAllCount()
	if err != nil {
		return nil, err
	}

	return &PaginatedModels{
		Result: langChainLLMs,
		Total:  int(langChainLLMsCount),
	}, nil
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
