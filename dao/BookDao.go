package dao

import (
	. "../config"
	. "../model"
)

type IBookDao interface {
	Create(book *Book)
}

type BookDaoImpl struct {}

var BookDao IBookDao = (*BookDaoImpl)(nil)

func (BookDaoImpl) Create(book *Book) {
	GormDB.Create(book)
}