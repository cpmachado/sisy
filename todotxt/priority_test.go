package todotxt

import "testing"

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want string
		p    *Priority
	}{
		{name: "A", want: "(A)", p: NewPriority('A')},
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
		{name: "A", prio: 'A', want: NewPriority('A')},
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
			if want.Cmp(got) != 0 {
				t.Errorf("NewPriority() = %v, want %v", got, want)
			}
		})
	}
}

func TestPriority_Cmp(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		a    rune
		b    rune
		want int
	}{
		{name: "a = b", a: 'A', b: 'A', want: 0},
		{name: "a > b", a: 'A', b: 'B', want: 1},
		{name: "a < b", a: 'B', b: 'A', want: -1},
		{name: "a is nil/invalid", a: '0', b: 'B', want: -26},
		{name: "b is nil/invalid", a: 'A', b: '0', want: 26},
		{name: "a, b are nil/invalid", a: '0', b: '0', want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			a := NewPriority(tt.a)
			b := NewPriority(tt.b)
			want := tt.want

			// then
			got := a.Cmp(b)

			if got != want {
				t.Errorf("Cmp() = %v, want %v", got, want)
			}
		})
	}
}
