package todotxt

import (
	"strings"
	"time"
)

// Todo definition
type Todo struct {
	Complete       bool
	Priority       *Priority
	CompletionDate *time.Time
	CreationDate   *time.Time
	Description    Description
}

func (t *Todo) String() string {
	sb := &strings.Builder{}
	if t.Complete {
		sb.WriteString("x ")
	}

	if t.Priority != nil {
		sb.WriteString(t.Priority.String())
		sb.WriteString(" ")
	}

	if t.CompletionDate != nil {
		sb.WriteString(t.CompletionDate.Format(time.RFC3339))
		sb.WriteString(" ")
	}

	if t.CreationDate != nil {
		sb.WriteString(t.CompletionDate.Format(time.RFC3339))
		sb.WriteString(" ")
	}

	sb.WriteString(t.Description.String())

	return sb.String()
}
