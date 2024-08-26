package models

import (
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/third_party/langchain"
)

type Repo struct {
	db *database.Database
}

type Service struct {
	repo      *Repo
	langchain *langchain.LangChain
}
