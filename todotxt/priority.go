package todotxt

import "fmt"

// priority definition
type Priority rune

func (p *Priority) String() string {
	return fmt.Sprintf("(%c)", *p)
}
