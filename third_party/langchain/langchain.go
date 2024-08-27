package langchain

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/third_party/logger"
)

type LangChainToken = string

type LangChain struct {
	logger *logger.Logger
	store  *Store
	db     *database.Database
}

func New(logger *logger.Logger, db *database.Database) *LangChain {
	return &LangChain{
		store:  NewStore(logger),
		logger: logger,
		db:     db,
	}
}

func (lc *LangChain) Prompt(token LangChainToken, prompt PromptPayload) error {
	llm, err := lc.store.GetLLM(token)
	if err != nil {
		return err
	}

	return llm.Prompt(prompt)
}

func (lc *LangChain) AddLLM(params AddLLMParams) error {
	storedLLM := models.LangChainLLM{
		UserID:   params.UserID,
		Token:    params.Token,
		Model:    params.Model,
		Provider: params.Provider,
		URL:      params.URL,
		APIKey:   params.APIKey,
		Status:   "pending",
		Enabled:  true,
	}

	err := lc.db.Create(&storedLLM).Error
	if err != nil {
		return err
	}

	err = lc.store.AddLLM(params)
	if err != nil {
		storedLLM.Status = "error"
		storedLLM.Enabled = false
		storedLLM.Error = err.Error()
		err := lc.db.Save(&storedLLM).Error
		if err != nil {
			lc.logger.Error().Err(err).Msg("failed to update stored LLM")
			return err
		}
	}

	storedLLM.Status = "ok"
	storedLLM.Enabled = true
	storedLLM.Error = ""
	err = lc.db.Save(&storedLLM).Error
	if err != nil {
		lc.logger.Error().Err(err).Msg("failed to update stored LLM")
		return err
	}

	return nil
}

func (lc *LangChain) RemoveLLM(token LangChainToken) error {
	storedLLM := new(models.LangChainLLM)
	err := lc.db.Find(storedLLM).
		Where("token = ?", token).
		Error
	if err != nil {
		return err
	}

	err = lc.db.Delete(storedLLM).Error
	if err != nil {
		return err
	}

	err = lc.store.RemoveLLM(token)
	if err != nil {
		return err
	}

	return nil
}

func (lc *LangChain) InitializeLLMs() error {
	lc.logger.Info().Msg("initializing LLMs")
	storedLLMs := make([]models.LangChainLLM, 0)
	err := lc.db.Find(&storedLLMs).
		Where("enabled = ?", true).
		Error
	if err != nil {
		return err
	}

	for _, storedLLM := range storedLLMs {
		lc.logger.Info().Msgf("initializing LLM: %s", storedLLM.Token)
		err = lc.store.AddLLM(AddLLMParams{
			Token:    storedLLM.Token,
			Model:    storedLLM.Model,
			Provider: storedLLM.Provider,
			URL:      storedLLM.URL,
			APIKey:   storedLLM.APIKey,
		})

		if err != nil {
			storedLLM.Status = "error"
			storedLLM.Enabled = false
			storedLLM.Error = err.Error()
			err := lc.db.Save(storedLLM).Error
			if err != nil {
				lc.logger.Error().Err(err).Msg("failed to update stored LLM")
			}
			continue
		}

		storedLLM.Status = "ok"
		storedLLM.Enabled = true
		storedLLM.Error = ""
		err := lc.db.Save(storedLLM).Error
		if err != nil {
			lc.logger.Error().Err(err).Msg("failed to update stored LLM")
		}
	}

	lc.logger.Info().Msg("LLMs initialized")
	return nil
}
