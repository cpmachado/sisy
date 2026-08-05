package todotxt

import (
	"fmt"
	"unicode"
)

// priority definition, uppercase letters, less is more
type Priority rune

const priorityMaxD = int('Z' - 'A')

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
		return int(*b) - int(*p) // less is more
	case p != nil:
		return priorityMaxD + 1
	case b != nil:
		return -priorityMaxD - 1
	}
	return 0
}
