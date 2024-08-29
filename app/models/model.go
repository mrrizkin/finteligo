package models

import "gorm.io/gorm"

type Seed interface {
	Seed(db *gorm.DB)
}

type Model struct {
	models []interface{}
	seeds  []Seed
}

func New() *Model {
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
			&User{},
		},
	}
}

func (m *Model) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m.models...)
}

func (m *Model) Seeds(db *gorm.DB) error {
	for _, model := range m.seeds {
		model.Seed(db)
	}

	return nil
}
