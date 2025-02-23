package todo

import (
	"bufio"
	"os"
)

// SaveTodos writes the todos list to a Todo.txt file
func SaveTodos(todos []Todo, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, todo := range todos {
		var line string

		// If the todo is completed, mark it with 'x'
		if todo.Completed {
			line = "x "
		}

		// Add the priority if it exists
		if todo.Priority != "" {
			line += "(" + todo.Priority + ") "
		}

		// Add the project labels if they exist
		for _, project := range todo.Projects {
			line += "+" + project + " "
		}

		// Add the context labels if they exist
		for _, context := range todo.Contexts {
			line += "@" + context + " "
		}

		// Add the due date if it exists
		if !todo.DueDate.IsZero() {
			line += "due:" + todo.DueDate.Format("2006-01-02") + " "
		}

		// Add the title (task description)
		line += todo.Title

		// Write the line to the file
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
