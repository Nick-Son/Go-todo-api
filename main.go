package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// creates a new router
	router := mux.NewRouter().StrictSlash(true)
	// adds a "/" route, assigns the Index handler when the endpoint is called
	router.HandleFunc("/", Index)
	// spin up a server on port 8080, and can be accessed at http://localhost:8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
