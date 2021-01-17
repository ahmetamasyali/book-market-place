package main

import (
	. "./config"
	_ "./controller"
	"log"
	"net/http"
)


func main() {
	InitDb()

	log.Fatal(http.ListenAndServe(":8080", nil))
}




