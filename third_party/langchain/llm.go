package langchain

import (
	"context"
	"errors"

	"github.com/tmc/langchaingo/llms"
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
	Role        string  `json:"role"`
	Content     string  `json:"content"`
	Model       string  `json:"model"       validate:"required"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        int     `json:"top_k"`
	Message     string  `json:"message" validate:"required"`
	StreamFunc  *func(ctx context.Context, chunk []byte) error
	Channel     chan string
}

func NewLLM(provider string, token LangChainToken) *LLM {
	return &LLM{
		Provider: provider,
		Token:    token,
	}
}

func (llm *LLM) Prompt(prompt PromptPayload) error {
	switch llm.Provider {
	case "ollama":
		return llm.promptOllama(prompt)
	case "anthropic":
		return llm.promptAnthropic(prompt)
	case "openai":
		return llm.promptOpenai(prompt)
	default:
		return ErrInvalidProvider
	}
}

func (llm *LLM) promptOllama(prompt PromptPayload) error {
	ctx := context.Background()

	options := make([]llms.CallOption, 0)

	if prompt.Temperature != 0 {
		options = append(options, llms.WithTemperature(prompt.Temperature))
	}

	if prompt.TopP != 0 {
		options = append(options, llms.WithTopP(prompt.TopP))
	}

	if prompt.TopK != 0 {
		options = append(options, llms.WithTopK(prompt.TopK))
	}

	if prompt.StreamFunc != nil {
		options = append(options, llms.WithStreamingFunc(*prompt.StreamFunc))
	}

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		llm.ollamaLLM,
		prompt.Message,
		options...,
	)
	if err != nil {
		return err
	}

	prompt.Channel <- completion
	return nil
}

func (llm *LLM) promptAnthropic(prompt PromptPayload) error {
	return nil
}

func (llm *LLM) promptOpenai(prompt PromptPayload) error {
	return nil
}
