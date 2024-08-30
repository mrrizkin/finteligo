package ollama

import (
	"context"

	"github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
	lcOllama "github.com/tmc/langchaingo/llms/ollama"
)

type Ollama struct {
	llm *lcOllama.LLM
}

func New(params types.AddLLMParams) (types.LLM, error) {
	config := make([]lcOllama.Option, 0)

	if params.URL != "" {
		config = append(config, lcOllama.WithServerURL(params.URL))
	}

	if params.Model != "" {
		config = append(config, lcOllama.WithModel(params.Model))
	}

	llm, err := lcOllama.New(config...)
	if err != nil {
		return nil, err
	}

	return &Ollama{
		llm: llm,
	}, nil
}

func (o *Ollama) ChatPrompt(payload types.PromptPayload) error {
	ctx := context.Background()

	options := make([]llms.CallOption, 0)

	if payload.Temperature != 0 {
		options = append(options, llms.WithTemperature(payload.Temperature))
	}

	if payload.TopP != 0 {
		options = append(options, llms.WithTopP(payload.TopP))
	}

	if payload.TopK != 0 {
		options = append(options, llms.WithTopK(payload.TopK))
	}

	if payload.StreamFunc != nil {
		options = append(options, llms.WithStreamingFunc(*payload.StreamFunc))
	}

	completion, err := o.llm.GenerateContent(ctx, payload.Messages, options...)
	if err != nil {
		return err
	}

	payload.Channel <- completion.Choices[0].Content
	return nil
}

func (o *Ollama) Prompt(payload types.PromptPayload) error {
	ctx := context.Background()

	options := make([]llms.CallOption, 0)

	if payload.Temperature != 0 {
		options = append(options, llms.WithTemperature(payload.Temperature))
	}

	if payload.TopP != 0 {
		options = append(options, llms.WithTopP(payload.TopP))
	}

	if payload.TopK != 0 {
		options = append(options, llms.WithTopK(payload.TopK))
	}

	if payload.StreamFunc != nil {
		options = append(options, llms.WithStreamingFunc(*payload.StreamFunc))
	}

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		o.llm,
		payload.Message,
		options...,
	)
	if err != nil {
		return err
	}

	payload.Channel <- completion
	return nil
}

func (o *Ollama) SinglePrompt(payload types.PromptPayload) error {
	ctx := context.Background()

	options := make([]llms.CallOption, 0)

	if payload.Temperature != 0 {
		options = append(options, llms.WithTemperature(payload.Temperature))
	}

	if payload.TopP != 0 {
		options = append(options, llms.WithTopP(payload.TopP))
	}

	if payload.TopK != 0 {
		options = append(options, llms.WithTopK(payload.TopK))
	}

	if payload.StreamFunc != nil {
		options = append(options, llms.WithStreamingFunc(*payload.StreamFunc))
	}

	completion, err := llms.GenerateFromSinglePrompt(
		ctx,
		o.llm,
		payload.Message,
		options...,
	)
	if err != nil {
		return err
	}

	payload.Channel <- completion
	return nil
}
