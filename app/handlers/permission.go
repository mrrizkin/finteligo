package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) PermissionCreate(c *fiber.Ctx) error {
	payload := new(models.Permission)
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

	permission, err := h.permissionService.Create(payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed create permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed create permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create permission",
		Data:    permission,
	})
}

func (h *Handlers) PermissionFindAll(c *fiber.Ctx) error {
	var (
		err         error
		permissions []models.Permission
	)

	permissions, err = h.permissionService.FindAll()
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get permissions")
		return &fiber.Error{
			Code:    500,
			Message: "failed get permissions",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get permissions",
		Data:    permissions,
	})
}

func (h *Handlers) PermissionFindByID(c *fiber.Ctx) error {
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

	permission, err := h.permissionService.FindByID(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed get permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed get permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get permission",
		Data:    permission,
	})
}

func (h *Handlers) PermissionUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed to parse id")
		return &fiber.Error{
			Code:    400,
			Message: "id not valid",
		}
	}

	payload := new(models.Permission)
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

	permission, err := h.permissionService.Update(uint(id), payload)
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed update permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed update permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update permission",
		Data:    permission,
	})
}

func (h *Handlers) PermissionDelete(c *fiber.Ctx) error {
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

	err = h.permissionService.Delete(uint(id))
	if err != nil {
		h.System.Logger.Error().Err(err).Msg("failed delete permission")
		return &fiber.Error{
			Code:    500,
			Message: "failed delete permission",
		}
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete permission",
	})
}
