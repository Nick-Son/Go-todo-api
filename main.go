package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

// Todo - define a type called Todo which is a struct(Go's version of a class)
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

// Todos - defined a Todos type that is an array of items of type Todo
type Todos []Todo

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
	// fmt.Fprintln(w, "Todo Index")
	todos := Todos{
		Todo{Name: "Create Github repo"},
		Todo{Name: "Push code to Github"},
	}

	// serve todos up as a JSON response
	json.NewEncoder(w).Encode(todos)
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
