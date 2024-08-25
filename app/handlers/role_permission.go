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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	rolePermission, err := h.rolePermissionService.Create(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create role permission",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get role permissions",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	rolePermission, err := h.rolePermissionService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get role permission",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	payload := new(models.RolePermission)
	err = c.BodyParser(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	rolePermission, err := h.rolePermissionService.Update(uint(id), payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update role permission",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = h.rolePermissionService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete role permission",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete role permission",
	})
}
