package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mrrizkin/finteligo/app/config"
	"github.com/mrrizkin/finteligo/resources"
	"github.com/mrrizkin/finteligo/system/session"
	"github.com/mrrizkin/finteligo/system/validator"
	"github.com/mrrizkin/finteligo/third_party/logger"
)

type Server struct {
	*fiber.App

	config *config.Config
}

type Site struct {
	Menus []MenuItem
}

type MenuItem struct {
	Name string
	URL  string
	Icon string
}

func New(config *config.Config, logger *logger.Logger, session *session.Session) *Server {
	app := fiber.New(fiber.Config{
		Prefork:               config.PREFORK,
		AppName:               config.APP_NAME,
		DisableStartupMessage: true,
		Views:                 html.NewFileSystem(http.FS(resources.ViewFS), ".html"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(validator.GlobalErrorResponse{
				Status:  "error",
				Title:   http.StatusText(code),
				Detail:  err.Error(),
				Message: err.Error(),
			})
		},
	})

	if config.ENV != "development" {
		csrfConfig := csrf.Config{
			KeyLookup:         "header:" + csrf.HeaderName,
			CookieName:        "finteligo_csrf_token",
			CookieSameSite:    "Lax",
			CookieSecure:      false,
			CookieSessionOnly: true,
			CookieHTTPOnly:    true,
			SingleUseToken:    true,
			Expiration:        1 * time.Hour,
			KeyGenerator:      utils.UUIDv4,
			ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
			Extractor:         csrf.CsrfFromHeader(csrf.HeaderName),
			Session:           session.Store,
			SessionKey:        "fiber.csrf.token",
			HandlerContextKey: "fiber.csrf.handler",
		}

		app.Use(csrf.New(csrfConfig))
	}

	app.Static("/", "public")

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.Logger,
	}))
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(idempotency.New())

	return &Server{
		App:    app,
		config: config,
	}
}

func (s *Server) Serve() error {
	return s.Listen(fmt.Sprintf(":%d", s.config.PORT))
}

func (s *Server) Stop() error {
	return s.Shutdown()
}
