package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mrrizkin/finteligo/app/config"
	"github.com/mrrizkin/finteligo/resources"
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

func New(config *config.Config, logger *logger.Logger) *Server {
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
