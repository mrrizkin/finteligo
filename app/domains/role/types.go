package role

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

type PaginatedRole struct {
	Result []models.Role
	Total  int
}

type RolePayload struct {
	models.Role
	PermissionIDs []uint `json:"permission_ids"`
}
