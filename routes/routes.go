package routes

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/mrrizkin/finteligo/app/handlers"
	"github.com/mrrizkin/finteligo/system/session"
	"github.com/mrrizkin/finteligo/system/types"
)

func Setup(app *types.App, session *session.Session) {
	handler := handlers.New(app)
	api := app.Group("/api")
	ApiRoutes(api, handler)

	csrfConfig := csrf.Config{
		KeyLookup:         "cookie:" + csrf.HeaderName,
		CookieName:        "finteligo_csrf_token",
		CookieSameSite:    "Lax",
		CookieSecure:      false,
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
		SingleUseToken:    true,
		Expiration:        1 * time.Hour,
		KeyGenerator:      utils.UUIDv4,
		ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
		Extractor:         csrf.CsrfFromCookie("finteligo_csrf_token"),
		Session:           session.Store,
		SessionKey:        "fiber.csrf.token",
		HandlerContextKey: "fiber.csrf.handler",
	}
	app.Use(csrf.New(csrfConfig))
	WebRoutes(app, handler)
}
