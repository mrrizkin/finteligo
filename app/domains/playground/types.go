package playground

import (
	"github.com/mrrizkin/finteligo/third_party/langchain"
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
)

type Service struct {
	langchain *langchain.LangChain
}

type PromptPayload struct {
	types.PromptPayload
	Token types.Token `json:"token"`
}
