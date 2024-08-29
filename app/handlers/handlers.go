package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/domains/api_token"
	"github.com/mrrizkin/finteligo/app/domains/auth"
	"github.com/mrrizkin/finteligo/app/domains/models"
	"github.com/mrrizkin/finteligo/app/domains/permission"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/app/domains/role"
	"github.com/mrrizkin/finteligo/app/domains/role_permission"
	"github.com/mrrizkin/finteligo/app/domains/user"
	dbModels "github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

type Handlers struct {
	*types.App

	apiTokenRepo    *api_token.Repo
	apiTokenService *api_token.Service

	authRepo    *auth.Repo
	authService *auth.Service

	modelsRepo    *models.Repo
	modelsService *models.Service

	permissionRepo    *permission.Repo
	permissionService *permission.Service

	playgroundService *playground.Service

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
	apiTokenRepo := api_token.NewRepo(app.System.Database)
	apiTokenService := api_token.NewService(apiTokenRepo)

	authRepo := auth.NewRepo(app.System.Database)
	authService := auth.NewService(authRepo)

	modelsRepo := models.NewRepo(app.System.Database)
	modelsService := models.NewService(modelsRepo, app.Library.LangChain)

	permissionRepo := permission.NewRepo(app.System.Database)
	permissionService := permission.NewService(permissionRepo)

	playgroundService := playground.NewService(app.Library.LangChain)

	rolePermissionRepo := role_permission.NewRepo(app.System.Database)
	rolePermissionService := role_permission.NewService(rolePermissionRepo)

	roleRepo := role.NewRepo(app.System.Database)
	roleService := role.NewService(roleRepo)

	userRepo := user.NewRepo(app.System.Database)
	userService := user.NewService(userRepo)

	return &Handlers{
		App: app,

		apiTokenRepo:    apiTokenRepo,
		apiTokenService: apiTokenService,

		authRepo:    authRepo,
		authService: authService,

		modelsRepo:    modelsRepo,
		modelsService: modelsService,

		permissionRepo:    permissionRepo,
		permissionService: permissionService,

		playgroundService: playgroundService,

		rolePermissionRepo:    rolePermissionRepo,
		rolePermissionService: rolePermissionService,

		roleRepo:    roleRepo,
		roleService: roleService,

		userRepo:    userRepo,
		userService: userService,
	}
}

func (h *Handlers) GetUser(c *fiber.Ctx) *dbModels.User {
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

func (h *Handlers) GetPaginationQuery(c *fiber.Ctx) types.Pagination {
	page := c.QueryInt("page", 1)
	perPage := c.QueryInt("per_page", 10)

	return types.Pagination{
		Page:    page,
		PerPage: perPage,
	}
}
