package langchain

import (
	"errors"

	"github.com/mrrizkin/finteligo/third_party/logger"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	ErrLLMNotFound = errors.New("llm not found")
)

type Store struct {
	logger *logger.Logger
	llms   map[LangChainToken]*LLM
}

func NewStore(logger *logger.Logger) *Store {
	return &Store{
		llms:   make(map[string]*LLM),
		logger: logger,
	}
}

type AddLLMParams struct {
	UserID   uint
	Token    LangChainToken
	Model    string
	Provider string
	URL      string
	APIKey   string
}

func (s *Store) GetLLM(token LangChainToken) (*LLM, error) {
	if llm, ok := s.llms[token]; ok {
		return llm, nil
	}

	return nil, ErrLLMNotFound
}

func (s *Store) AddLLM(params AddLLMParams) error {
	llm := NewLLM(params.Provider, params.Token)

	switch params.Provider {
	case "ollama":
		ollamaLLM, err := s.createOllamaLLM(params)
		if err != nil {
			return err
		}

		llm.ollamaLLM = ollamaLLM
	case "anthropic":
		anthropicLLM, err := s.createAnthropicLLM(params)
		if err != nil {
			return err
		}

		llm.anthropicLLM = anthropicLLM
	case "openai":
		openaiLLM, err := s.createOpenaiLLM(params)
		if err != nil {
			return err
		}

		llm.openaiLLM = openaiLLM
	default:
		return errors.New("unknown provider")
	}

	s.llms[params.Token] = llm

	return nil
}

func (s *Store) RemoveLLM(token LangChainToken) error {
	if _, ok := s.llms[token]; ok {
		delete(s.llms, token)
		return nil
	}

	return ErrLLMNotFound
}

func (s *Store) createOllamaLLM(params AddLLMParams) (*ollama.LLM, error) {
	config := make([]ollama.Option, 0)

	if params.URL != "" {
		config = append(config, ollama.WithServerURL(params.URL))
	}

	if params.Model != "" {
		config = append(config, ollama.WithModel(params.Model))
	}

	llm, err := ollama.New(config...)
	if err != nil {
		return nil, err
	}
	return llm, nil
}

func (s *Store) createAnthropicLLM(params AddLLMParams) (*anthropic.LLM, error) {
	config := make([]anthropic.Option, 0)

	if params.URL != "" {
		config = append(config, anthropic.WithBaseURL(params.URL))
	}

	if params.APIKey != "" {
		config = append(config, anthropic.WithToken(params.APIKey))
	}

	if params.Model != "" {
		config = append(config, anthropic.WithModel(params.Model))
	}

	llm, err := anthropic.New(config...)
	if err != nil {
		return nil, err
	}

	return llm, nil
}

func (s *Store) createOpenaiLLM(params AddLLMParams) (*openai.LLM, error) {
	config := make([]openai.Option, 0)

	if params.URL != "" {
		config = append(config, openai.WithBaseURL(params.URL))
	}

	if params.APIKey != "" {
		config = append(config, openai.WithToken(params.APIKey))
	}

	if params.Model != "" {
		config = append(config, openai.WithModel(params.Model))
	}

	llm, err := openai.New(config...)
	if err != nil {
		return nil, err
	}

	return llm, nil
}

func (s *Store) logInfo(msg string) {
	if s.logger != nil {
		s.logger.Info().Msg("LLM Store: " + msg)
	}
}

func (s *Store) logError(msg string) {
	if s.logger != nil {
		s.logger.Error().Msg("LLM Store: " + msg)
	}
}

func (s *Store) logDebug(msg string) {
	if s.logger != nil {
		s.logger.Debug().Msg("LLM Store: " + msg)
	}
}
