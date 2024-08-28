package langchain

import (
	"errors"

	"github.com/mrrizkin/finteligo/third_party/langchain/provider/anthropic"
	"github.com/mrrizkin/finteligo/third_party/langchain/provider/ollama"
	"github.com/mrrizkin/finteligo/third_party/langchain/provider/openai"
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
)

var (
	ErrLLMNotFound = errors.New("llm not found")
)

type Store struct {
	llms map[types.Token]types.LLM
}

func NewStore() *Store {
	return &Store{
		llms: make(map[string]types.LLM),
	}
}

func (s *Store) GetLLM(token types.Token) (types.LLM, error) {
	if llm, ok := s.llms[token]; ok {
		return llm, nil
	}

	return nil, ErrLLMNotFound
}

func (s *Store) AddLLM(params types.AddLLMParams) error {
	switch params.Provider {
	case "ollama":
		llm, err := ollama.New(params)
		if err != nil {
			return err
		}

		s.llms[params.Token] = llm
		return nil
	case "anthropic":
		llm, err := anthropic.New(params)
		if err != nil {
			return err
		}

		s.llms[params.Token] = llm
		return nil
	case "openai":
		llm, err := openai.New(params)
		if err != nil {
			return err
		}

		s.llms[params.Token] = llm
		return nil
	default:
		return errors.New("unknown provider")
	}
}

func (s *Store) RemoveLLM(token types.Token) error {
	if _, ok := s.llms[token]; ok {
		delete(s.llms, token)
		return nil
	}

	return ErrLLMNotFound
}
