package seed

import (
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

// All function seeds database with dummy data that is necessary for tests
func All() []Seed {
	return []Seed{
		{
			Name: "Example-entity-1",
			Run: func(db *gorm.DB) error {
				return createExampleEntity(db, 1, 2)
			},
		},
	}
}
