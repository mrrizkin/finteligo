package routes

import (
	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/system/types"
)

func WebRoutes(app *types.App) {
	handler := handlers.New(app)

	ui := app.Group("/_")
	ui.Get("/mimin", handler.AdminUI)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/playground/prompt", handler.Prompt)

	v1.Get("/permission", handler.PermissionFindAll)
	v1.Get("/permission/:id", handler.PermissionFindByID)
	v1.Post("/permission", handler.PermissionCreate)
	v1.Put("/permission/:id", handler.PermissionUpdate)
	v1.Delete("/permission/:id", handler.PermissionDelete)

	v1.Get("/role_permission", handler.RolePermissionFindAll)
	v1.Get("/role_permission/:id", handler.RolePermissionFindByID)
	v1.Post("/role_permission", handler.RolePermissionCreate)
	v1.Put("/role_permission/:id", handler.RolePermissionUpdate)
	v1.Delete("/role_permission/:id", handler.RolePermissionDelete)

	v1.Get("/role", handler.RoleFindAll)
	v1.Get("/role/:id", handler.RoleFindByID)
	v1.Post("/role", handler.RoleCreate)
	v1.Put("/role/:id", handler.RoleUpdate)
	v1.Delete("/role/:id", handler.RoleDelete)

	v1.Get("/user", handler.UserFindAll)
	v1.Get("/user/:id", handler.UserFindByID)
	v1.Post("/user", handler.UserCreate)
	v1.Put("/user/:id", handler.UserUpdate)
	v1.Delete("/user/:id", handler.UserDelete)
}
