package routes

import (
	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/system/types"
)

func Setup(app *types.App) {
	handler := handlers.New(app)
	api := app.Group("/api")
	ApiRoutes(api, handler)
	WebRoutes(app, handler)
}
