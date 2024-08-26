package langchain

import (
	"errors"

	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	ErrInvalidProvider = errors.New("invalid provider")
)

type LLM struct {
	Provider     string
	Token        LangChainToken
	ollamaLLM    *ollama.LLM
	anthropicLLM *anthropic.LLM
	openaiLLM    *openai.LLM
}

type PromptPayload struct {
	Role    string `json:"role"`
	Content string `json:"content"`

	Model       string  `json:"model"       validate:"required"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        int     `json:"top_k"`

	Message string `json:"message" validate:"required"`
}

func NewLLM(provider string, token LangChainToken) *LLM {
	return &LLM{
		Provider: provider,
		Token:    token,
	}
}

func (llm *LLM) Prompt(prompt PromptPayload) (string, error) {
	switch llm.Provider {
	case "ollama":
		return llm.promptOllama(prompt)
	case "anthropic":
		return llm.promptAnthropic(prompt)
	case "openai":
		return llm.promptOpenai(prompt)
	default:
		return "", ErrInvalidProvider
	}
}

func (llm *LLM) promptOllama(prompt PromptPayload) (string, error) {
	return "", nil
}

func (llm *LLM) promptAnthropic(prompt PromptPayload) (string, error) {
	return "", nil
}

func (llm *LLM) promptOpenai(prompt PromptPayload) (string, error) {
	return "", nil
}
