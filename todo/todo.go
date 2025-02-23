package todo

import "time"

// Todo represents a single Todo item, adapted for Todo.txt format
type Todo struct {
	Title      string    // The task description
	Priority   string    // The task priority, e.g., "(A)"
	Projects   []string  // The projects associated with the task, e.g., "+work", "+home"
	Contexts   []string  // The contexts associated with the task, e.g., "@home", "@office"
	DueDate    time.Time // The due date of the task, e.g., "due:2025-12-25"
	HasDueDate bool      // hasDueDate
	Completed  bool      // The completion status of the task
}
