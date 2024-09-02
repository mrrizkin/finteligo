package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) ModelsCreate(c *fiber.Ctx) error {
	payload := new(models.LangChainLLM)
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

	if payload.UserID == 0 {
		payload.UserID = 1
	}

	models, err := h.modelsService.Create(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create models")
		return &fiber.Error{
			Code:    500,
			Message: "failed create models",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success create models",
		Data:    models,
	})
}

func (h *Handlers) ModelsFindAll(c *fiber.Ctx) error {
	pagination := h.GetPaginationQuery(c)
	modelss, err := h.modelsService.FindAll(pagination)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get modelss")
		return &fiber.Error{
			Code:    500,
			Message: "failed get modelss",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get modelss",
		Data:    modelss.Result,
		Meta: &types.PaginationMeta{
			Page:      pagination.Page,
			PerPage:   pagination.PerPage,
			Total:     modelss.Total,
			PageCount: modelss.Total / pagination.PerPage,
		},
	})
}

func (h *Handlers) ModelsFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	models, err := h.modelsService.FindByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return &fiber.Error{
				Code:    404,
				Message: "models not found",
			}
		}
		h.System.Logger.Error().Err(err).Msg("failed get models")
		return &fiber.Error{
			Code:    500,
			Message: "failed get models",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get models",
		Data:    models,
	})
}

func (h *Handlers) ModelsUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.LangChainLLM)
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

	models, err := h.modelsService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update models")
		return &fiber.Error{
			Code:    500,
			Message: "failed update models",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success update models",
		Data:    models,
	})
}

func (h *Handlers) ModelsDelete(c *fiber.Ctx) error {
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

	err = h.modelsService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete models")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete models",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success delete models",
	})
}
