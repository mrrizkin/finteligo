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

func (s *Service) Prompt(
	token types.Token,
	payload types.PromptPayload,
) error {
	err := s.langchain.Prompt(token, payload)
	if err != nil {
		return err
	}

	return nil
}
