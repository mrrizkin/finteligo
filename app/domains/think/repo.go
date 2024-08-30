package think

import (
	"github.com/mrrizkin/finteligo/app/domains/think/templates"
	"github.com/mrrizkin/finteligo/app/domains/think/types"
	"github.com/tmc/langchaingo/llms"
)

type Repo struct {
	templates map[string]types.PromptTemplate
}

func NewRepo() *Repo {
	templates := map[string]types.PromptTemplate{
		"ppatp": templates.NewPPATP(),
	}

	return &Repo{
		templates: templates,
	}
}

func (r *Repo) Create(useCase, prompt string) []llms.MessageContent {
	return r.templates[useCase].GenContent(llms.TextParts(llms.ChatMessageTypeGeneric, prompt))
}

func (r *Repo) Generate(useCase string, payload interface{}) string {
	return r.templates[useCase].GenMessage(payload)
}

func (r *Repo) Parse(output string) (interface{}, error) {
	return r.templates["ppatp"].OutputParser(output)
}
