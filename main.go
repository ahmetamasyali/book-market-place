package main

import (
	. "./config"
	. "./service"
	"log"
	"net/http"
)


func main() {
	InitDb()

	http.HandleFunc("/createPerson", PersonService.CreatePersonHandler)
	http.HandleFunc("/createBook", BookService.CreateBookHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}




