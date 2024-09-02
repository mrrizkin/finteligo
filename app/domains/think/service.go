package think

import (
	"github.com/mrrizkin/finteligo/app/domains/think/types"
	"github.com/mrrizkin/finteligo/third_party/langchain"
	lcTypes "github.com/mrrizkin/finteligo/third_party/langchain/types"
)

type Service struct {
	langchain *langchain.LangChain
	repo      *Repo
}

func NewService(repo *Repo, langchain *langchain.LangChain) *Service {
	return &Service{
		repo:      repo,
		langchain: langchain,
	}
}

func (s *Service) AskAI(useCase string, payload *types.PromptPayload) error {
	messages := s.repo.Create(useCase, payload.Message)

	promptPayload := lcTypes.PromptPayload{
		Role:        payload.Role,
		Content:     payload.Content,
		Temperature: payload.Temperature,
		TopP:        payload.TopP,
		TopK:        payload.TopK,
		Message:     payload.Message,
		StreamFunc:  payload.StreamFunc,
		Channel:     payload.Channel,
		Messages:    messages,
	}

	err := s.langchain.ChatPrompt(payload.Token, promptPayload)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GenerateMessage(useCase string, payload interface{}) string {
	return s.repo.Generate(useCase, payload)
}

func (s *Service) OutputParser(output string) (interface{}, error) {
	return s.repo.Parse(output)
}
