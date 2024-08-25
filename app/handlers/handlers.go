package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/domains/permission"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/app/domains/role"
	"github.com/mrrizkin/finteligo/app/domains/role_permission"
	"github.com/mrrizkin/finteligo/app/domains/user"
	"github.com/mrrizkin/finteligo/system/types"
)

type Handlers struct {
	*types.App

	playgroundService *playground.Service

	permissionRepo    *permission.Repo
	permissionService *permission.Service

	rolePermissionRepo    *role_permission.Repo
	rolePermissionService *role_permission.Service

	roleRepo    *role.Repo
	roleService *role.Service

	userRepo    *user.Repo
	userService *user.Service
}

func New(
	app *types.App,
) *Handlers {

	playgroundService := playground.NewService(&playground.Llm{})

	permissionRepo := permission.NewRepo(app.System.Database)
	permissionService := permission.NewService(permissionRepo)

	rolePermissionRepo := role_permission.NewRepo(app.System.Database)
	rolePermissionService := role_permission.NewService(rolePermissionRepo)

	roleRepo := role.NewRepo(app.System.Database)
	roleService := role.NewService(roleRepo)

	userRepo := user.NewRepo(app.System.Database)
	userService := user.NewService(userRepo)

	return &Handlers{
		App: app,

		playgroundService: playgroundService,

		permissionRepo:    permissionRepo,
		permissionService: permissionService,

		rolePermissionRepo:    rolePermissionRepo,
		rolePermissionService: rolePermissionService,

		roleRepo:    roleRepo,
		roleService: roleService,

		userRepo:    userRepo,
		userService: userService,
	}
}

func (h *Handlers) SendJson(c *fiber.Ctx, resp interface{}, status ...int) error {
	var statusCode int

	if len(status) == 0 {
		statusCode = 200
	} else {
		statusCode = status[0]
	}

	return c.Status(statusCode).JSON(resp)
}
