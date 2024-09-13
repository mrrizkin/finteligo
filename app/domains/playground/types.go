package playground

import (
	"github.com/mrrizkin/finteligo/third_party/langchain"
	"github.com/mrrizkin/finteligo/third_party/langchain/types"
)

type Service struct {
	langchain *langchain.LangChain
}

type ChatHistory struct {
	Role    string   `json:"role"`
	Content []string `json:"content"`
}

type PromptPayload struct {
	types.PromptPayload
	ChatHistory []ChatHistory `json:"chat_history"`
	Token       types.Token   `json:"token"`
}
