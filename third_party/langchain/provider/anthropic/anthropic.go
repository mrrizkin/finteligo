package anthropic

import (
	"context"

	"github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
	lcAnthropic "github.com/tmc/langchaingo/llms/anthropic"
)

type Anthropic struct {
	llm *lcAnthropic.LLM
}

func New(params types.AddLLMParams) (types.LLM, error) {
	config := make([]lcAnthropic.Option, 0)

	if params.URL != "" {
		config = append(config, lcAnthropic.WithBaseURL(params.URL))
	}

	if params.APIKey != "" {
		config = append(config, lcAnthropic.WithToken(params.APIKey))
	}

	if params.Model != "" {
		config = append(config, lcAnthropic.WithModel(params.Model))
	}

	llm, err := lcAnthropic.New(config...)
	if err != nil {
		return nil, err
	}

	return &Anthropic{
		llm: llm,
	}, nil
}

func (o *Anthropic) Prompt(payload types.PromptPayload) error {
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

func (o *Anthropic) SinglePrompt(payload types.PromptPayload) error {
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
