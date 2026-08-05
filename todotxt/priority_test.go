package todotxt

import "testing"

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
		p    *Priority
	}{
		{name: "a", want: "(a)", p: NewPriority('a')},
		{name: "invalid", want: "", p: NewPriority('1')},
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

func TestNewPriority(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		prio rune
		want *Priority
	}{
		{name: "a", prio: 'a', want: NewPriority('a')},
		{name: "invalid", prio: '2', want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			prio := tt.prio
			want := tt.want

			// then
			got := NewPriority(prio)

			// when
			if !want.Equal(got) {
				t.Errorf("NewPriority() = %v, want %v", got, want)
			}
		})
	}
}
