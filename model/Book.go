package model


type Book struct {
	Id uint `gorm:"PRIMARY_KEY;NOT NULL"`
	Name string `gorm:"NOT NULL"`
	PersonId int
}

func (Book) TableName() string {
	return "book"
}

type NewBookRequest struct {
	Name string
	PersonId int
}
