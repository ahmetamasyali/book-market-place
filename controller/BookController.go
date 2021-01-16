package controller

import (
	. "../model"
	. "../service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type IBookController interface {
	CreateBookHandler(http.ResponseWriter, *http.Request)
}

type BookControllerImpl struct {}

var BookController IBookController = (*BookControllerImpl)(nil)

func (BookControllerImpl) CreateBookHandler(responseWriter http.ResponseWriter, req *http.Request) {
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

	var person = PersonService.LoadById(newBookRequest.PersonId)

	if person == nil {
		http.Error(responseWriter, fmt.Sprintf("person with id %d not fount", newBookRequest.PersonId), http.StatusBadRequest)
	}

	BookService.Create(&Book{Name: newBookRequest.Name, PersonId: newBookRequest.PersonId})
}