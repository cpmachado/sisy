package main

import (
	"fmt"
	"iter"
	"slices"
	"time"

	"go.cpmachado.pt/sisy"
)

type simpleTodo struct {
	content string
}

func (s *simpleTodo) Content() string {
	return s.content
}

func (s *simpleTodo) CreationDate() *time.Time {
	t := time.Now()
	return &t
}

func (s *simpleTodo) CompletionDate() *time.Time {
	panic("not implemented") // TODO: Implement
}

func (s *simpleTodo) Status() string {
	return "TODO"
}

func main() {
	todos := []sisy.Todo{
		&simpleTodo{
			content: "Call mom",
		},
		&simpleTodo{
			content: "Buy Groceries",
		},
		&simpleTodo{
			content: "Read something",
		},
		&simpleTodo{
			content: "Dread about the world",
		},
	}

	listTodos(slices.Values(todos))
}

func listTodos(todos iter.Seq[sisy.Todo]) {
	fmt.Println("Just listing:")

	for todo := range todos {
		fmt.Printf("- %s: %q | %s\n", todo.CreationDate().Format(time.DateTime), todo.Status(), todo.Content())
	}
}
