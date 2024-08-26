package playground

import "github.com/mrrizkin/finteligo/third_party/langchain"

func NewService(lc *langchain.LangChain) *Service {
	return &Service{
		langchain: lc,
	}
}

func (s *Service) Prompt(
	token langchain.LangChainToken,
	payload langchain.PromptPayload,
) (*PromptResponse, error) {
	result, err := s.langchain.Prompt(token, payload)
	if err != nil {
		return nil, err
	}

	return &PromptResponse{
		Answer: result,
	}, nil
}
