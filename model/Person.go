package model

type Person struct {
	Id uint `gorm:"PRIMARY_KEY;NOT NULL"`
	Name string `gorm:"NOT NULL"`
	Books []Book `gorm:"foreignKey:PersonId"`
}

func (Person) TableName() string {
	return "person"
}

type NewPersonRequest struct {
	Name string
}