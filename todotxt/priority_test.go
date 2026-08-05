package todotxt

import (
	"testing"
)

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
		p    *Priority
	}{
		{name: "a", want: "(a)", p: NewPriority('a')},
		{name: "nil", want: "", p: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			p := tt.p
			want := tt.want

			// when
			got := p.String()

			// then
			if got != want {
				t.Errorf("String() = %v, want %v", got, want)
			}
		})
	}
}
