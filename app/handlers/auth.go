package handlers

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/system/types"
)

type LoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handlers) Identity(c *fiber.Ctx) error {
	return h.SendJson(c, types.Response{
		Title:   "Identity",
		Message: "You are authorized",
		Status:  "success",
	})
}

func (h *Handlers) Login(c *fiber.Ctx) error {
	payload := new(LoginPayload)
	if err := c.BodyParser(payload); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid payload",
		}
	}

	validationError := h.System.Validator.MustValidate(payload)
	if validationError != nil {
		return validationError
	}

	user, err := h.authService.Login(payload.Username, payload.Password)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Invalid username or password",
		}
	}

	session, err := h.System.Session.Get(c)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to get session",
		}
	}

	sid := session.ID()
	uid := user.ID
	session.Set("uid", uid)
	session.Set("sid", sid)
	session.Set("ip", c.Context().RemoteIP().String())
	session.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
	session.Set("ua", string(c.Request().Header.UserAgent()))

	err = session.Save()
	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "Failed to save session",
		}
	}

	return h.SendJson(c, types.Response{
		Title:   "Login successfully",
		Message: "Welcome back",
		Status:  "success",
	})
}

type LogoutPayload struct {
	SID string `json:"sid"`
}

func (h *Handlers) Logout(c *fiber.Ctx) error {
	req := new(LogoutPayload)
	if err := c.BodyParser(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid payload",
		}
	}

	// Get current session
	session, _ := h.System.Session.Get(c)

	// Check session ID
	if len(req.SID) > 0 {
		// Get requested session
		data, err := h.System.Session.Storage.Get(req.SID)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Failed to get session",
			}
		}

		// Decode requested session data
		gd := gob.NewDecoder(bytes.NewBuffer(data))
		dm := make(map[string]interface{})
		if err := gd.Decode(&dm); err != nil {
			return &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Failed to decode session",
			}
		}

		// If it belongs to current user destroy requested session
		if session.Get("uid").(string) == dm["uid"] {
			h.System.Session.Storage.Delete(req.SID)
		}
	} else {
		// Destroy current session
		session.Destroy()
	}

	return h.SendJson(c, types.Response{
		Title:   "Logout successfully",
		Message: "Have a nice day",
		Status:  "success",
	})
}
