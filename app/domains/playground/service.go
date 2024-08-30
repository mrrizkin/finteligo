package playground

import (
	"github.com/mrrizkin/finteligo/third_party/langchain"
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
)

func NewService(lc *langchain.LangChain) *Service {
	return &Service{
		langchain: lc,
	}
}

func (s *Service) Prompt(payload *PromptPayload) error {
	promptPayload := types.PromptPayload{
		Role:        payload.Role,
		Content:     payload.Content,
		Model:       payload.Model,
		Temperature: payload.Temperature,
		TopP:        payload.TopP,
		TopK:        payload.TopK,
		Message:     payload.Message,
		StreamFunc:  payload.StreamFunc,
		Channel:     payload.Channel,
	}

	err := s.langchain.Prompt(payload.Token, promptPayload)
	if err != nil {
		return err
	}

	return nil
}
