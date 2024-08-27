package routes

import (
	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/system/types"
)

func WebRoutes(app *types.App, handler *handlers.Handlers) {
	ui := app.Group("/_")
	ui.Get("/mimin", handler.AdminUI)
}
