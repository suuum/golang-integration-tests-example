package entities

type ExampleEntity struct {
	ID         int
	FistProp   int `gorm:"column:first_prop"`
	SecondProp int `gorm:"column:second_prop"`
}

func (ExampleEntity) TableName() string {
	return "example_entities"
}
