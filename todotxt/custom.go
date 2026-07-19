package todotxt

import "fmt"

// extension parser definition
type CustomTag struct {
	Key   string
	Value string
}

func (c *CustomTag) String() string {
	return fmt.Sprintf("%s:%s", c.Key, c.Value)
}
