package sisy

import "time"

type Todo interface {
	Content() string
	CreationDate() *time.Time
	CompletionDate() *time.Time
	Status() string
}
