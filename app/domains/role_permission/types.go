package role_permission

import "github.com/mrrizkin/finteligo/system/database"

type Service struct {
	repo *Repo
}

type Repo struct {
	db *database.Database
}
