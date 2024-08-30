package handlers

import (
	"github.com/mrrizkin/finteligo/app/domains/api_token"
	"github.com/mrrizkin/finteligo/app/domains/auth"
	"github.com/mrrizkin/finteligo/app/domains/models"
	"github.com/mrrizkin/finteligo/app/domains/permission"
	"github.com/mrrizkin/finteligo/app/domains/playground"
	"github.com/mrrizkin/finteligo/app/domains/role"
	"github.com/mrrizkin/finteligo/app/domains/role_permission"
	"github.com/mrrizkin/finteligo/app/domains/think"
	"github.com/mrrizkin/finteligo/app/domains/user"
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

	thinkRepo    *think.Repo
	thinkService *think.Service

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

	thinkRepo := think.NewRepo()
	thinkService := think.NewService(thinkRepo, app.Library.LangChain)

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

		thinkRepo:    thinkRepo,
		thinkService: thinkService,

		roleRepo:    roleRepo,
		roleService: roleService,

		userRepo:    userRepo,
		userService: userService,
	}
}
