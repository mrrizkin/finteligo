package models

import (
	"gorm.io/gorm"

	"github.com/mrrizkin/finteligo/app/config"
	"github.com/mrrizkin/finteligo/third_party/argon2"
)

type Seed interface {
	Seed(db *gorm.DB)
}

type Model struct {
	argon2 *argon2.Argon2
	config *config.Config
	models []interface{}
	seeds  []Seed
}

func New(config *config.Config, argon2 *argon2.Argon2) *Model {
	return &Model{
		models: []interface{}{
			&Permission{},
			&Plan{},
			&RolePermission{},
			&Role{},
			&Status{},
			&Subscription{},
			&User{},
			&LangChainLLM{},
			&ApiToken{},
		},
		seeds: []Seed{
			&Permission{},
			&Role{},
			&RolePermission{},
			&Status{},
		},
		argon2: argon2,
		config: config,
	}
}

func (m *Model) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m.models...)
}

func (m *Model) Seeds(db *gorm.DB) error {
	for _, model := range m.seeds {
		model.Seed(db)
	}

	userModel := new(User)
	userModel.Seed(m.config, m.argon2, db)

	return nil
}
