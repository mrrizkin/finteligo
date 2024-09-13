package playground

import (
	"strings"

	"github.com/mrrizkin/finteligo/third_party/langchain"
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
)

func NewService(lc *langchain.LangChain) *Service {
	return &Service{
		langchain: lc,
	}
}

func (s *Service) Prompt(payload *PromptPayload) error {
	Messages := make([]llms.MessageContent, 0)

	for _, message := range payload.ChatHistory {
		if message.Role == "assistant" {
			content := strings.Join(message.Content, "\n")
			Messages = append(Messages, llms.TextParts(llms.ChatMessageTypeAI, content))
		} else if message.Role == "user" {
			content := strings.Join(message.Content, "\n")
			Messages = append(Messages, llms.TextParts(llms.ChatMessageTypeHuman, content))
		}
	}

	Messages = append(Messages, llms.TextParts(llms.ChatMessageTypeHuman, payload.Message))

	promptPayload := types.PromptPayload{
		Role:        payload.Role,
		Content:     payload.Content,
		Temperature: payload.Temperature,
		TopP:        payload.TopP,
		TopK:        payload.TopK,
		Messages:    Messages,
		StreamFunc:  payload.StreamFunc,
		Channel:     payload.Channel,
	}

	// err := s.langchain.Prompt(payload.Token, promptPayload)
	// if err != nil {
	// 	return err
	// }

	err := s.langchain.ChatPrompt(payload.Token, promptPayload)
	if err != nil {
		return err
	}

	return nil
}
