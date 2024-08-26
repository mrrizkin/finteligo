package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) RolePermissionCreate(c *fiber.Ctx) error {
	payload := new(models.RolePermission)
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

	rolePermission, err := h.rolePermissionService.Create(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create role permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed create role permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create role permission",
		Data:    rolePermission,
	})
}

func (h *Handlers) RolePermissionFindAll(c *fiber.Ctx) error {
	rolePermission, err := h.rolePermissionService.FindAll()
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get role permissions")
		return &fiber.Error{
			Code:    500,
			Message: "failed get role permissions",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get role permissions",
		Data:    rolePermission,
	})
}

func (h *Handlers) RolePermissionFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	rolePermission, err := h.rolePermissionService.FindByID(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get role permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed get role permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get role permission",
		Data:    rolePermission,
	})
}

func (h *Handlers) RolePermissionUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.RolePermission)
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

	rolePermission, err := h.rolePermissionService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update role permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed update role permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update role permission",
		Data:    rolePermission,
	})
}

func (h *Handlers) RolePermissionDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	err = h.rolePermissionService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete role permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete role permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete role permission",
	})
}
