package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"os"
	"text/template"

	"go.cpmachado.pt/sisy/todo"
)

//go:embed template/*
var templates embed.FS

var (
	todoFile string = "todos.txt"
	todos    []todo.Todo
	tmpl     *template.Template
)

func init() {
	flag.StringVar(&todoFile, "f", todoFile, "Todo file used")
	flag.Parse()
	var err error
	todos, err = LoadTodos(todoFile)
	if err != nil {
		fmt.Printf("Error loading todos: %v\n", err)
		return
	}
	// Load the template from a file
	tmpl, err = template.ParseFS(templates, "template/todos.tmpl", "template/todo.tmpl")
	if err != nil {
		fmt.Printf("Error loading template: %v\n", err)
		return
	}
}

func main() {
	PrintTodos(todos, tmpl)
}

// LoadTodos reads todos from a Todo.txt file
func LoadTodos(filename string) ([]todo.Todo, error) {
	var todos []todo.Todo
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		todo, err := todo.ParseTodo(line)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// PrintTodos prints all todos to the console using a template
func PrintTodos(todos []todo.Todo, tmpl *template.Template) {
	err := tmpl.Execute(os.Stdout, todos)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
}
