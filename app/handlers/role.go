package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) RoleCreate(c *fiber.Ctx) error {
	payload := new(models.Role)
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

	role, err := h.roleService.Create(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create role")
		return &fiber.Error{
			Code:    500,
			Message: "failed create role",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create role",
		Data:    role,
	})
}

func (h *Handlers) RoleFindAll(c *fiber.Ctx) error {
	roles, err := h.roleService.FindAll()
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get roles")
		return &fiber.Error{
			Code:    500,
			Message: "failed get roles",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get roles",
		Data:    roles,
	})
}

func (h *Handlers) RoleFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	role, err := h.roleService.FindByID(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get role")
		return &fiber.Error{
			Code:    500,
			Message: "failed get role",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get role",
		Data:    role,
	})
}

func (h *Handlers) RoleUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.Role)
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

	role, err := h.roleService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update role")
		return &fiber.Error{
			Code:    500,
			Message: "failed update role",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update role",
		Data:    role,
	})
}

func (h *Handlers) RoleDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	err = h.roleService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete role")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete role",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete role",
	})
}
