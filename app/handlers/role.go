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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	role, err := h.roleService.Create(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create role",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get roles",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	role, err := h.roleService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get role",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	payload := new(models.Role)
	err = c.BodyParser(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	role, err := h.roleService.Update(uint(id), payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update role",
			Debug:   err.Error(),
		}, 500)
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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = h.roleService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete role",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete role",
	})
}
