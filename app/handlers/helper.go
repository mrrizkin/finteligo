package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func (h *Handlers) GetUser(c *fiber.Ctx) *models.User {
	userId := c.Locals("uid").(uint)
	user, err := h.userRepo.FindByID(userId)
	if err != nil {
		return nil
	}

	return user
}

func (h *Handlers) SendJson(c *fiber.Ctx, resp types.Response, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}

func (h *Handlers) GetUserByToken(token string) (*models.User, error) {
	apiToken, err := h.apiTokenRepo.FindByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := h.userRepo.FindByID(apiToken.UserId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *Handlers) GetPaginationQuery(c *fiber.Ctx) types.Pagination {
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per_page", 10)

	return types.Pagination{
		Page:    page,
		PerPage: perPage,
	}
}
