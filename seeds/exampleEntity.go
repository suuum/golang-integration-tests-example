package seed

import (
	"example/entities"

	"gorm.io/gorm"
)

func createExampleEntity(db *gorm.DB, firstProp, secondProp int) error {
	return db.Create(&entities.ExampleEntity{
		FistProp:   firstProp,
		SecondProp: secondProp,
	}).Error
}
