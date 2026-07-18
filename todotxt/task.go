package todotxt

import "time"

// task definition
type Task struct {
	Complete       *bool
	Priority       *Priority
	CompletionDate *time.Time
	CreationDate   *time.Time
	Description    Description
}
