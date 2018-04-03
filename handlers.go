package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

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
