package main

import (
	. "./config"
	. "./controller"
	"log"
	"net/http"
)


func main() {
	InitDb()

	http.HandleFunc("/createPerson", PersonController.CreatePersonHandler)
	http.HandleFunc("/createBook", BookController.CreateBookHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}




