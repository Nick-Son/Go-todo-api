package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	// ./main.go:10:12: undefined: NewRouter
	log.Fatal(http.ListenAndServe(":8080", router))
}
