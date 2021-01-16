package dao

import (
	. "../config"
	. "../model"
	"strconv"
)

type IPersonDao interface {
	LoadById(id int) *Person
	Create(person *Person)
}

type PersonDaoImpl struct {}

var PersonDao IPersonDao = PersonDaoImpl{}


func (PersonDaoImpl) LoadById(id int) *Person {
	var person *Person
	GormDB.First(person, strconv.Itoa(id))

	return person
}

func (PersonDaoImpl) Create(person *Person) {
	GormDB.Create(person)
}


