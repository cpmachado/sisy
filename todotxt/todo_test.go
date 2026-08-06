package todotxt

import (
	"testing"
	"time"
)

func aDate(year int, month time.Month, day int) *time.Time {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return &t
}

func TestTodo_String(t *testing.T) {
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

func TestTodo_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    []byte
		want    *Todo
		wantErr bool
	}{
		{
			name: "call mom",
			text: []byte("x 2011-03-03 Call Mom"),
			want: &Todo{
				Complete:     true,
				CreationDate: aDate(2011, time.March, 3),
				Description: Description{
					Text: "Call Mom",
				},
			},
			wantErr: false,
		},
		{
			name: "thank mom",
			text: []byte("(A) Thank Mom for the meatballs @phone"),
			want: &Todo{
				Priority: NewPriority('A'),
				Description: Description{
					Text: "Thank Mom for the meatballs @phone",
				},
			},
			wantErr: false,
		},
		{
			name: "complete without priority",
			text: []byte("x 2011-03-02 2011-03-01 Review Tim's pull request +TodoTxtTouch @github"),
			want: &Todo{
				Complete:       true,
				CompletionDate: aDate(2011, time.March, 2),
				CreationDate:   aDate(2011, time.March, 1),
				Description: Description{
					Text: "Review Tim's pull request +TodoTxtTouch @github",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var to Todo
			gotErr := to.UnmarshalText(tt.text)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalText() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalText() succeeded unexpectedly")
			} else if to.String() != tt.want.String() {
				t.Errorf("UnmarshalText() = %q, want %q", &to, tt.want)
			}
		})
	}
}
