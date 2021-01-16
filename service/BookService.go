package service

import (
	. "../dao"
	. "../model"
)

type IBookService interface {
	Create(book *Book)
}

type BookServiceImpl struct {}

var BookService IBookService = BookServiceImpl{}

func (BookServiceImpl) Create(book *Book) {
	BookDao.Create(book)
}