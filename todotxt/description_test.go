package todotxt

import (
	"testing"
)

func TestDescription_String(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        string
		description Description
	}{
		{
			name: "base",
			want: "Thank Mom for the meatballs @phone",
			description: Description{
				Text:           "Thank Mom for the meatballs",
				ProjectContext: []string{"phone"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			description := tt.description
			want := tt.want

			// when
			got := description.String()

			if got != want {
				t.Errorf("String() = %v, want %v", got, want)
			}
		})
	}
}
