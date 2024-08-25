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
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := h.userService.Create(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create user",
		Data:    user,
	})
}

func (h *Handlers) UserFindAll(c *fiber.Ctx) error {
	users, err := h.userService.FindAll()
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get users",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get users",
		Data:    users,
	})
}

func (h *Handlers) UserFindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := h.userService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get user",
		Data:    user,
	})
}

func (h *Handlers) UserUpdate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	payload := new(models.User)
	err = c.BodyParser(payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := h.userService.Update(uint(id), payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update user",
		Data:    user,
	})
}

func (h *Handlers) UserDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = h.userService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete user",
	})
}
