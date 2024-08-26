package playground

import "github.com/mrrizkin/finteligo/third_party/langchain"

type Service struct {
	langchain *langchain.LangChain
}

type PromptPayload struct {
	langchain.PromptPayload
	Token langchain.LangChainToken `json:"token"`
}

type PromptResponse struct {
	Answer string `json:"answer"`
}
