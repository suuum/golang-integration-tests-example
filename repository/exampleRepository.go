package repository

import (
	"example/entities"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	GetEntity(firstParm, secondParam int) entities.ExampleEntity
}

func CreateExampleRepositoryRepository(db *gorm.DB) *ExampleRepositoryStruct {
	return &ExampleRepositoryStruct{DB: db}
}

func (repo *ExampleRepositoryStruct) GetEntity(firstParm, secondParam int) entities.ExampleEntity {
	var result *entities.ExampleEntity
	err := repo.DB.Model(entities.ExampleEntity{}).Where("first_prop = ? and second_prop = ?", firstParm, secondParam).First(&result).Error
	if err != nil {
		logrus.Error("query error", err)
	}

	return *result
}

type ExampleRepositoryStruct struct {
	DB *gorm.DB
}
