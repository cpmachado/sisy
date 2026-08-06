package todotxt

import (
	"strings"
	"time"
	"unicode"
)

// Todo definition
type Todo struct {
	Complete       bool
	Priority       *Priority
	CompletionDate *time.Time
	CreationDate   *time.Time
	Description    Description
}

func (t *Todo) MarshalText() (text []byte, err error) {
	if t.Priority != nil && !t.Priority.IsValid() {
		return nil, ErrInvalidPriority
	}

	if t.CompletionDate != nil && t.CreationDate == nil {
		return nil, ErrMissingCreationDate
	}

	return []byte(t.String()), nil
}

func (t *Todo) UnmarshalText(text []byte) error {
	const (
		completed int = iota
		priority
		dates
		description
	)
	states := []int{
		completed, priority, dates, description,
	}

	nextNonSpace := func(base int, text []byte) int {
		for i, c := range text[base:] {
			if !unicode.IsSpace(rune(c)) {
				return base + i
			}
		}
		return len(text)
	}

	ptr := 0

	for _, state := range states {
		switch state {
		case completed:
			if text[ptr] == 'x' {
				t.Complete = true
				ptr++
			}
		case priority:
			if text[ptr] == '(' {
				if r := rune(text[ptr+1]); unicode.IsUpper(r) && text[ptr+2] == ')' {
					t.Priority = NewPriority(r)
					if t.Priority != nil {
						ptr += 3
					}
				}
			}
		case dates:
			subbuf := text[ptr:]
			if len(subbuf) > 10 {
				tcomp, err := time.Parse(time.DateOnly, string(subbuf[:10]))
				if err != nil {
					break
				}
				nptr := nextNonSpace(10, subbuf)
				subbuf = subbuf[nptr:]
				tcreat, err := time.Parse(time.DateOnly, string(subbuf[:10]))
				if err != nil {
					t.CreationDate = &tcomp
					ptr += 10
				} else {
					t.CompletionDate = &tcomp
					t.CreationDate = &tcreat
					ptr += nptr + 10
				}
			}
		case description:
			t.Description.Init(string(text[ptr:]))
		}
		ptr = nextNonSpace(ptr, text)
	}

	return nil
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
		sb.WriteString(t.CompletionDate.Format(time.DateOnly))
		sb.WriteString(" ")
	}

	if t.CreationDate != nil {
		sb.WriteString(t.CreationDate.Format(time.DateOnly))
		sb.WriteString(" ")
	}

	sb.WriteString(t.Description.String())

	return sb.String()
}

func (t *Todo) ProjectTags() []string {
	if t == nil {
		return nil
	}
	return t.Description.ProjectTags()
}

func (t *Todo) ContextTags() []string {
	if t == nil {
		return nil
	}
	return t.Description.ContextTags()
}

func (t *Todo) Special() map[string]string {
	if t == nil {
		return nil
	}
	return t.Description.Special()
}
