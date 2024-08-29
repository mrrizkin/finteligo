package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) ApiTokenCreate(c *fiber.Ctx) error {
	payload := new(models.ApiToken)
	err := c.BodyParser(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse payload")
		return &fiber.Error{
			Code:    400,
			Message: "payload not valid",
		}
	}
	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	user := h.GetUser(c)
	if user == nil {
		return &fiber.Error{
			Code:    401,
			Message: "Unauthorized",
		}
	}

	apiToken, err := h.apiTokenService.Create(payload, user)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create apiToken")
		return &fiber.Error{
			Code:    500,
			Message: "failed create api token",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success create api token",
		Data:    apiToken,
	})
}

func (h *Handlers) ApiTokenFindAll(c *fiber.Ctx) error {
	pagination := h.GetPaginationQuery(c)
	apiTokens, err := h.apiTokenService.FindAll(pagination)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get apiTokens")
		return &fiber.Error{
			Code:    500,
			Message: "failed get api tokens",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get api tokens",
		Data:    apiTokens.Result,
		Meta: types.PaginationMeta{
			Page:      pagination.Page,
			PerPage:   pagination.PerPage,
			Total:     apiTokens.Total,
			PageCount: apiTokens.Total / pagination.PerPage,
		},
	})
}

func (h *Handlers) ApiTokenFindByID(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	apiToken, err := h.apiTokenService.FindByID(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get api token")
		return &fiber.Error{
			Code:    500,
			Message: "failed get api token",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get api token",
		Data:    apiToken,
	})
}

func (h *Handlers) ApiTokenUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.ApiToken)
	err = c.BodyParser(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse payload")
		return &fiber.Error{
			Message: "payload not valid",
			Code:    400,
		}
	}
	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	apiToken, err := h.apiTokenService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update api token")
		return &fiber.Error{
			Code:    500,
			Message: "failed update api token",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success update api token",
		Data:    apiToken,
	})
}

func (h *Handlers) ApiTokenDelete(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	err = h.apiTokenService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete api token")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete api token",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success delete api token",
	})
}
