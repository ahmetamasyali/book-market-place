package service

import (
	. "../dao"
	. "../model"
)

type IPersonService interface {
	LoadById(id int) *Person
	Create(person *Person)
}

type PersonServiceImpl struct {}

var PersonService IPersonService = PersonServiceImpl{}

func (PersonServiceImpl) LoadById(id int) *Person {
	return PersonDao.LoadById(id)
}

func (PersonServiceImpl) Create(person *Person) {
	PersonDao.Create(person)
}


