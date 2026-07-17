package sisy

import "time"

type Task struct {
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Dependencies []Task     `json:"dependencies"`
	Deadline     *time.Time `json:"deadline"`
	Scheduled    *TimeBlock `json:"scheduled"`
	Categories   []Category `json:"categories"`
	Tags         []Tag      `json:"tags"`
}

type TimeBlock struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
}

type Category string

type Tag string
