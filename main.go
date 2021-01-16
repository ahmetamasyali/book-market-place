package main

import (
	. "./config"
	. "./controller"
	_ "./dao"
	. "./service"
	"log"
	"net/http"
)


func main() {
	InitDb()

	PersonService.LoadById(1)
	http.HandleFunc("/createPerson", PersonController.CreatePersonHandler)
	http.HandleFunc("/createBook", BookController.CreateBookHandler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}




