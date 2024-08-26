package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/system/types"
	"github.com/mrrizkin/finteligo/third_party/langchain"
)

func (h *Handlers) Prompt(c *fiber.Ctx) error {
	payload := new(playground.PromptPayload)
	err := c.BodyParser(payload)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "payload not valid",
		}
	}

	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	promptResponse, err := h.playgroundService.Prompt(
		payload.Token,
		langchain.PromptPayload{
			Role:        payload.Role,
			Content:     payload.Content,
			Model:       payload.Model,
			Temperature: payload.Temperature,
			TopP:        payload.TopP,
			TopK:        payload.TopK,
			Message:     payload.Message,
		},
	)

	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: "failed prompt",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success prompt",
		Data:    promptResponse,
	})
}
