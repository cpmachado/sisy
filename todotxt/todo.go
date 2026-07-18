package todotxt

import "time"

// Todo definition
type Todo struct {
	Complete       *bool
	Priority       *Priority
	CompletionDate *time.Time
	CreationDate   *time.Time
	Description    Description
}
