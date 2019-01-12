package main

import "fmt"

var (
	todos Todos
	currentID int
)

func init()  {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meet up"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos  {
		if t.ID == id {
			todos = append(todos[:i], todos[i + i:]...)
			return nil
		}
	}

	return fmt.Errorf("could not find todo with id of %d to delete", id)
}