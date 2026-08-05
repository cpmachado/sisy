package todotxt

import (
	"fmt"
	"unicode"
)

// priority definition
type Priority rune

func NewPriority(prio rune) *Priority {
	if !unicode.IsUpper(prio) {
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
	return p != nil && unicode.IsUpper(rune(*p))
}

func (p *Priority) Cmp(b *Priority) int {
	switch {
	case p != nil && b != nil:
		return int(*p) - int(*b)
	case p != nil:
		return 1
	case b != nil:
		return -1
	}
	return 0
}
