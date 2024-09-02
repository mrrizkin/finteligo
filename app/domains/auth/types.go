package auth

import (
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/third_party/argon2"
)

type Repo struct {
	db *database.Database
}

type Service struct {
	repo   *Repo
	argon2 *argon2.Argon2
}
