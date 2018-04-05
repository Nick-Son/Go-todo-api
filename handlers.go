package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex handler
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// setting the conent type, telling the client to expect JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// setting the status code
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
	// serve todos up as a JSON response
	// json.NewEncoder(w).Encode(todos)
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

// TodoCreate handler
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	// open and read the content of body, limiting the amount of data to 1048576 bytes
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	// unmarshal the body to our todo struct
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
