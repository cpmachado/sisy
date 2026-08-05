package todotxt

import (
	"testing"
	"time"
)

func TestTodo_String(t *testing.T) {
	aDate := func(year int, month time.Month, day int) *time.Time {
		t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		return &t
	}
	tests := []struct {
		name string // description of this test case
		want string
		todo *Todo
	}{
		{
			name: "measure space",
			want: "x (A) 2016-05-20 2016-04-30 measure space for +chapelShelving @chapel due:2016-05-30",
			todo: &Todo{
				Complete:       true,
				Priority:       NewPriority('A'),
				CompletionDate: aDate(2016, time.May, 20),
				CreationDate:   aDate(2016, time.April, 30),
				Description: Description{
					Text: "measure space for +chapelShelving @chapel due:2016-05-30",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			to := tt.todo
			want := tt.want

			// when
			got := to.String()

			if got != want {
				t.Errorf("String() = %q, want %q", got, want)
			}
		})
	}
}
