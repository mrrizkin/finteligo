package api_token

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
)

type Repo struct {
	db *database.Database
}

type Service struct {
	repo *Repo
}

type PaginatedApiToken struct {
	Result []models.ApiToken
	Total  int
}
