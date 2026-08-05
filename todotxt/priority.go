package todotxt

import (
	"fmt"
	"unicode"
)

// priority definition
type Priority rune

func NewPriority(prio rune) *Priority {
	if !unicode.IsLetter(prio) {
		return nil
	}
	p := new(Priority)
	*p = Priority(prio)
	return p
}

func (p *Priority) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("(%c)", *p)
}
