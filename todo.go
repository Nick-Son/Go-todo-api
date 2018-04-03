package main

import "time"

// Todo - define a type called Todo which is a struct(Go's version of a class)
type Todo struct {
	// struct tags control how your struct will be marshalled to JSON.
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos - defined a Todos type that is an array of items of type Todo
type Todos []Todo
