package routes

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/system/types"
)

func WebRoutes(app *types.App, handler *handlers.Handlers) {

	ui := app.Group("/_", cors.New())
	ui.Get("/mimin", handler.AdminUI)
}
