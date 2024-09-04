package types

import (
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
	"github.com/tmc/langchaingo/llms"
)

type PromptTemplate interface {
	GenContent(content ...llms.MessageContent) []llms.MessageContent
	GenMessage(payload interface{}) string
	OutputParser(output string) (interface{}, error)
}

type PromptPayload struct {
	types.PromptPayload
	Token  types.Token `json:"token"`
	Stream bool        `json:"stream"`
}
