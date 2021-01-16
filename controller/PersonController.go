package controller

import (
	. "../model"
	. "../service"
	"encoding/json"
	"log"
	"net/http"
)

type IPersonController interface {
	CreatePersonHandler(http.ResponseWriter, *http.Request)
}

type PersonControllerImpl struct {}

var PersonController IPersonController = PersonControllerImpl{}

func (PersonControllerImpl) CreatePersonHandler(responseWriter http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(responseWriter,"", http.StatusNotFound)
		return
	}

	var newPersonRequest NewPersonRequest

	err := json.NewDecoder(req.Body).Decode(&newPersonRequest)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	if  len(newPersonRequest.Name) == 0 {
		log.Printf("person name cannot be empty")
		http.Error(responseWriter,"person name cannot be empty", http.StatusBadRequest)
		return
	}

	PersonService.Create(&Person{ Name: newPersonRequest.Name})

}