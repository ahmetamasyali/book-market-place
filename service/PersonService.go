package service

import (
	. "../config"
	. "../model"
	"encoding/json"
	"log"
	"net/http"
)

type IPersonService interface {
	CreatePersonHandler(http.ResponseWriter, *http.Request)
}

type PersonServiceImpl struct {}

var PersonService IPersonService = (*PersonServiceImpl)(nil)

func (PersonServiceImpl) CreatePersonHandler(responseWriter http.ResponseWriter, req *http.Request) {
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

	GormDB.Create(&Person{ Name: newPersonRequest.Name})

}