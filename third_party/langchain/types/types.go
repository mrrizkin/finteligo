package types

import (
	"context"

	"github.com/tmc/langchaingo/llms"
)

type LLM interface {
	Prompt(payload PromptPayload) error
	SinglePrompt(payload PromptPayload) error
	ChatPrompt(payload PromptPayload) error
}

type Token = string

type AddLLMParams struct {
	UserID   uint
	Token    Token
	Model    string
	Provider string
	URL      string
	APIKey   string
}

type PromptPayload struct {
	Role        string  `json:"role"`
	Content     string  `json:"content"`
	Model       string  `json:"model"       validate:"required"`
	Temperature float64 `json:"temperature"`
	TopP        float64 `json:"top_p"`
	TopK        int     `json:"top_k"`
	Message     string  `json:"message"     validate:"required"`
	Stream      bool    `json:"stream"`
	StreamFunc  *func(ctx context.Context, chunk []byte) error
	Channel     chan string
	Messages    []llms.MessageContent
}
