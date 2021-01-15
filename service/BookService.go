package service

import (
	. "../config"
	. "../model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type IBookService interface {
	CreateBookHandler(http.ResponseWriter, *http.Request)
}

type BookServiceImpl struct {}

var BookService IBookService = (*BookServiceImpl)(nil)

func (BookServiceImpl) CreateBookHandler(responseWriter http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(responseWriter, "", http.StatusNotFound)
		return
	}

	var newBookRequest NewBookRequest

	err := json.NewDecoder(req.Body).Decode(&newBookRequest)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	if newBookRequest.PersonId <= 0 {
		log.Printf("person id cannot be empty")
		http.Error(responseWriter, "person id be empty", http.StatusBadRequest)
		return
	}

	if len(newBookRequest.Name) == 0 {
		log.Printf("book name cannot be empty")
		http.Error(responseWriter, "book name cannot be empty", http.StatusBadRequest)
		return
	}

	var person Person
	GormDB.First(&person, strconv.Itoa(newBookRequest.PersonId))

	if person.Id <= 0 {
		http.Error(responseWriter, fmt.Sprintf("person with id %d not fount", newBookRequest.PersonId), http.StatusBadRequest)
	}

	GormDB.Create(&Book{Name: newBookRequest.Name, PersonId: newBookRequest.PersonId})
}