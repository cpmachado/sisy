package sisy

import "time"

// a definition of a task
type Task struct {
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Dependencies []Task     `json:"dependencies"`
	Deadline     *time.Time `json:"deadline"`
	Scheduled    *TimeBlock `json:"scheduled"`
	Categories   []Category `json:"categories"`
	Tags         []Tag      `json:"tags"`
}

// A definition of a recurring task
type RecurringTask struct {
	Current *Task `json:"current"`
	Next    func(Task) Task
}

// A definition of a time block
type TimeBlock struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
}

// a definition of a category
type Category string

// a definition of a tag
type Tag string
