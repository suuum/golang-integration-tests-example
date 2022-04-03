package database

import (
	"example/entities"

	"gorm.io/gorm"
)

func RunMigrations(gdb *gorm.DB) error {
	err := gdb.AutoMigrate(&entities.ExampleEntity{})

	return err
}
