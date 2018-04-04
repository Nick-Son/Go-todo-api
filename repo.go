package main

import "fmt"

var currentId int
var todos Todos

// initialise seed data
func init() {
	RepoCreateTodo(Todo{Name: "Feed Yuzu"})
	RepoCreateTodo(Todo{Name: "Clean Yuzu's Cage"})
}

// RepoFindTodo - find specific todo
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return an empty todo if not found
	return Todo{}
}

// RepoCreateTodo - create todo
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo - deletes a todo
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
