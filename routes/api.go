package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/routes/middleware"
)

func ApiRoutes(api fiber.Router, handler *handlers.Handlers) {
	api.Get("/identity", middleware.AuthProtected(handler.App, handler), handler.Identity)
	api.Post("/login", handler.Login)
	api.Post("/logout", handler.Logout)

	v1 := api.Group("/v1", middleware.AuthProtected(handler.App, handler))

	v1.Get("/api-tokens", handler.ApiTokenFindAll)
	v1.Get("/api-tokens/:id", handler.ApiTokenFindByID)
	v1.Post("/api-tokens", handler.ApiTokenCreate)
	v1.Post("/api-tokens/:id/disable", handler.ApiTokenDisable)
	v1.Post("/api-tokens/:id/enable", handler.ApiTokenEnable)
	v1.Put("/api-tokens/:id", handler.ApiTokenUpdate)
	v1.Delete("/api-tokens/:id", handler.ApiTokenDelete)

	v1.Get("/models", handler.ModelsFindAll)
	v1.Get("/models/:id", handler.ModelsFindByID)
	v1.Post("/models", handler.ModelsCreate)
	v1.Put("/models/:id", handler.ModelsUpdate)
	v1.Delete("/models/:id", handler.ModelsDelete)

	v1.Get("/permission", handler.PermissionFindAll)
	v1.Get("/permission/:id", handler.PermissionFindByID)
	v1.Post("/permission", handler.PermissionCreate)
	v1.Put("/permission/:id", handler.PermissionUpdate)
	v1.Delete("/permission/:id", handler.PermissionDelete)

	v1.Post("/playground/prompt", handler.Prompt)

	v1.Get("/role_permission", handler.RolePermissionFindAll)
	v1.Get("/role_permission/:id", handler.RolePermissionFindByID)
	v1.Post("/role_permission", handler.RolePermissionCreate)
	v1.Put("/role_permission/:id", handler.RolePermissionUpdate)
	v1.Delete("/role_permission/:id", handler.RolePermissionDelete)

	v1.Post("/ask-ai", handler.AskAI)

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
