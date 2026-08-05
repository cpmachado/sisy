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
	if !p.IsValid() {
		return ""
	}
	return fmt.Sprintf("(%c)", *p)
}

func (p *Priority) IsValid() bool {
	return p != nil && unicode.IsLetter(rune(*p))
}

func (p *Priority) Equal(b *Priority) bool {
	return (p != nil && b != nil && *p == *b) || p == b
}
