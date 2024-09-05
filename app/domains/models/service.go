package models

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/app/utils"
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
	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	err = s.langchain.AddLLM(lcTypes.AddLLMParams{
		UserID:   langChainLLM.UserID,
		Token:    token,
		Model:    langChainLLM.Model,
		Provider: langChainLLM.Provider,
		URL:      langChainLLM.URL,
		APIKey:   langChainLLM.APIKey,
	})
	if err != nil {
		return nil, err
	}

	lcLLM, err := s.repo.FindByToken(token)
	if err != nil {
		return nil, err
	}

	return lcLLM, nil
}

func (s *Service) FindAll(
	user *models.User,
	pagination types.Pagination,
) (*PaginatedModels, error) {
	wb := utils.NewWhereBuilder()

	wb.And("user_id = ?", user.ID)

	where, whereArgs := wb.Get()

	filter := types.Filter{
		Where:     where,
		WhereArgs: whereArgs,
	}

	langChainLLMs, err := s.repo.FindAll(pagination, filter)
	if err != nil {
		return nil, err
	}

	langChainLLMsCount, err := s.repo.FindAllCount(filter)
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

func generateToken() (string, error) {
	token, err := gonanoid.Generate(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-",
		32,
	)
	return "fin_ml_" + token, err
}
