package todotxt

import (
	"slices"
	"testing"
)

func TestDescription_String(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        string
		description Description
	}{
		{
			name: "one context",
			want: "Thank Mom for the meatballs @phone",
			description: Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: "Schedule Goodwill pickup +GarageSale @phone",
			description: Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: "Post signs around the neighborhood +GarageSale",
			description: Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: "@GroceryStore pies",
			description: Description{
				Text: "@GroceryStore pies",
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

func TestDescription_ProjectTags(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        []string
		description Description
	}{
		{
			name: "one context",
			want: nil,
			description: Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: []string{"GarageSale"},
			description: Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: []string{"GarageSale"},
			description: Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: nil,
			description: Description{
				Text: "@GroceryStore pies",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			description := tt.description
			want := tt.want

			// when
			got := description.ProjectTags()

			if !slices.Equal(got, want) {
				t.Errorf("ProjectTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
