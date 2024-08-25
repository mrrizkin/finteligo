package role

import "github.com/mrrizkin/finteligo/system/database"

type Repo struct {
	db *database.Database
}

type Service struct {
	repo *Repo
}
