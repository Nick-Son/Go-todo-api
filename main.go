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
	// routes for Todo index, Todo show and an about page
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/about", About)
	// spin up a server on port 8080, and can be accessed at http://localhost:8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex handler
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Todo Index")
}

// TodoShow handler
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Show Todos:", todoId)
}

// About handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple RESTful todo api built with Go :)")
}
