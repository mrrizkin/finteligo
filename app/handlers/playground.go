package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) Prompt(c *fiber.Ctx) error {
	payload := new(playground.Prompt)
	err := c.BodyParser(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	// pretty print payload
	fmt.Printf("%#+v\n", *payload)

	// promptResponse := h.playgroundService.Prompt(*payload)

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success prompt",
		// Data:    promptResponse,
	})
}
