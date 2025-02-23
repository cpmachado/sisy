package todo

import (
	"errors"
	"strings"
	"time"
)

// ParseTodo parses a single line of a Todo.txt file into a Todo struct
func ParseTodo(line string) (Todo, error) {
	todo := Todo{}
	line = strings.TrimSpace(line)

	// Check for completed task (starts with "x")
	if strings.HasPrefix(line, "x ") {
		todo.Completed = true
		line = strings.TrimPrefix(line, "x ")
	}

	// Parse the priority, e.g., "(A)"
	if strings.HasPrefix(line, "(") && strings.Contains(line, ")") {
		endIdx := strings.Index(line, ")")
		todo.Priority = line[1:endIdx]
		line = strings.TrimSpace(line[endIdx+1:])
	}

	// Parse multiple project labels, e.g., "+work +personal"
	for {
		projectIdx := strings.Index(line, "+")
		if projectIdx == -1 {
			break
		}
		projectEnd := strings.IndexAny(line[projectIdx:], " \t")
		if projectEnd == -1 {
			todo.Projects = append(todo.Projects, line[projectIdx+1:])
			break
		} else {
			todo.Projects = append(todo.Projects, line[projectIdx+1:projectIdx+projectEnd])
			line = line[projectIdx+projectEnd:]
		}
	}

	// Parse multiple context labels, e.g., "@home @work"
	for {
		contextIdx := strings.Index(line, "@")
		if contextIdx == -1 {
			break
		}
		contextEnd := strings.IndexAny(line[contextIdx:], " \t")
		if contextEnd == -1 {
			todo.Contexts = append(todo.Contexts, line[contextIdx+1:])
			break
		} else {
			todo.Contexts = append(todo.Contexts, line[contextIdx+1:contextIdx+contextEnd])
			line = line[contextIdx+contextEnd:]
		}
	}

	// Parse the due date, e.g., "due:YYYY-MM-DD"
	if strings.Contains(line, "due:") {
		dueDateIdx := strings.Index(line, "due:")
		dueDateStr := line[dueDateIdx+4:]
		todo.DueDate, _ = time.Parse("2006-01-02", dueDateStr)
		line = strings.Replace(line, "due:"+dueDateStr, "", 1)
		todo.HasDueDate = true
	}

	// The remaining part is the task description
	todo.Title = strings.TrimSpace(line)

	// If no title exists, return an error
	if todo.Title == "" {
		return Todo{}, errors.New("invalid todo item: missing task description")
	}

	return todo, nil
}
