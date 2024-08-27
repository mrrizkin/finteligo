package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) UserCreate(c *fiber.Ctx) error {
	payload := new(models.User)
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

	user, err := h.userService.Create(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create user")
		return &fiber.Error{
			Code:    500,
			Message: "failed create user",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success create user",
		Data:    user,
	})
}

func (h *Handlers) UserFindAll(c *fiber.Ctx) error {
	users, err := h.userService.FindAll()
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get users")
		return &fiber.Error{
			Code:    500,
			Message: "failed get users",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get users",
		Data:    users,
	})
}

func (h *Handlers) UserFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	user, err := h.userService.FindByID(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get user")
		return &fiber.Error{
			Code:    500,
			Message: "failed get user",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success get user",
		Data:    user,
	})
}

func (h *Handlers) UserUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.User)
	err = c.BodyParser(payload)
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

	user, err := h.userService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update user")
		return &fiber.Error{
			Code:    500,
			Message: "failed update user",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success update user",
		Data:    user,
	})
}

func (h *Handlers) UserDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	err = h.userService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete user")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete user",
		}
	}

	return h.SendJson(c, types.Response{
		Status:  "success",
		Title:   "Success",
		Message: "success delete user",
	})
}
