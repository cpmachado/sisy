package todotxt

import (
	"maps"
	"slices"
	"testing"
)

func TestDescription_String(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        string
		description *Description
	}{
		{
			name: "one context",
			want: "Thank Mom for the meatballs @phone",
			description: &Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: "Schedule Goodwill pickup +GarageSale @phone",
			description: &Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: "Post signs around the neighborhood +GarageSale",
			description: &Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: "@GroceryStore pies",
			description: &Description{
				Text: "@GroceryStore pies",
			},
		},
		{
			name:        "nil",
			want:        "",
			description: nil,
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
		description *Description
	}{
		{
			name: "one context",
			want: nil,
			description: &Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: []string{"GarageSale"},
			description: &Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: []string{"GarageSale"},
			description: &Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: nil,
			description: &Description{
				Text: "@GroceryStore pies",
			},
		},
		{
			name:        "nil",
			want:        nil,
			description: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			description := tt.description
			want := tt.want

			// when
			got := description.ProjectTags()

			// then
			if !slices.Equal(got, want) {
				t.Errorf("ProjectTags() = %v, want %v", got, want)
			}

			// test double parsing
			got2 := description.ProjectTags()

			if description != nil && (!slices.Equal(got, got2) || got != nil && &got[0] != &got2[0]) {
				t.Errorf("ProjectTags() is double parsing")
			}
		})
	}
}

func TestDescription_ContextTags(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        []string
		description *Description
	}{
		{
			name: "one context",
			want: []string{"phone"},
			description: &Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: []string{"phone"},
			description: &Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: nil,
			description: &Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: []string{"GroceryStore"},
			description: &Description{
				Text: "@GroceryStore pies",
			},
		},
		{
			name:        "nil",
			want:        nil,
			description: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			description := tt.description
			want := tt.want

			// when
			got := description.ContextTags()

			if !slices.Equal(got, want) {
				t.Errorf("ContextTags() = %v, want %v", got, want)
			}

			// test double parsing
			got2 := description.ContextTags()

			if description != nil && (!slices.Equal(got, got2) || got != nil && &got[0] != &got2[0]) {
				t.Errorf("ContextTags() is double parsing")
			}
		})
	}
}

func TestDescription_Special(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		want        map[string]string
		description *Description
	}{
		{
			name: "one context",
			want: nil,
			description: &Description{
				Text: "Thank Mom for the meatballs @phone",
			},
		},
		{
			name: "one tag one context",
			want: nil,
			description: &Description{
				Text: "Schedule Goodwill pickup +GarageSale @phone",
			},
		},
		{
			name: "one tag",
			want: nil,
			description: &Description{
				Text: "Post signs around the neighborhood +GarageSale",
			},
		},
		{
			name: "context interpolated with text",
			want: nil,
			description: &Description{
				Text: "@GroceryStore pies",
			},
		},
		{
			name: "context interpolated with text, with a special",
			want: map[string]string{"due": "2016-05-30"},
			description: &Description{
				Text: "@GroceryStore pies due:2016-05-30",
			},
		},
		{
			name:        "nil",
			want:        nil,
			description: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			description := tt.description
			want := tt.want

			// when
			got := description.Special()

			if !maps.Equal(got, want) {
				t.Errorf("Special() = %v, want %v", got, tt.want)
			}

			// test double parsing
			got2 := description.Special()

			if description != nil && !maps.Equal(got, got2) {
				t.Errorf("Special() is double parsing")
			}
		})
	}
}
