package playground

func NewService(llm *Llm) *Service {
	return &Service{llm}
}

func (s *Service) Prompt(payload Prompt) *PromptResponse {

	s.llm.Prompt(payload.Message)

	return &PromptResponse{
		Role:    payload.Role,
		Content: payload.Content,
	}
}
