package types

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/finteligo/app/config"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/system/session"
	"github.com/mrrizkin/finteligo/system/validator"
	"github.com/mrrizkin/finteligo/third_party/langchain"
	"github.com/mrrizkin/finteligo/third_party/logger"
)

type Response struct {
	Title   string         `json:"title"`
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Debug   string         `json:"debug"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type PaginationMeta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	PageCount int `json:"page_count"`
}

type App struct {
	*fiber.App

	System  *System
	Library *Library
}

type System struct {
	Logger    *logger.Logger
	Database  *database.Database
	Config    *config.Config
	Session   *session.Session
	Validator *validator.Validator
}

type Library struct {
	LangChain *langchain.LangChain
}
