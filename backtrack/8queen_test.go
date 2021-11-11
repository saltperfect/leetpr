package backtrack

import (
	"testing"
)

func Test_wrapper(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "first",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper()
		})
	}
}

func Test_checkIfGoodChoice(t *testing.T) {
	type args struct {
		grid [8][8]bool
		col  int
		row  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args {
				grid: [8][8]bool {
					{false, false, false ,false , false, false, false, false},
					{false, false, false ,false , false, false, false, true},
					{false, false, false ,false , false, false, false, true},
					{false, false, false ,false , false, false, false, true},
					{false, false, false ,false , false, false, false, true},
					{false, false, false ,true  , false, false, false, true},
					{false, false, false ,false , false, false, false, true},
					{false, false, false ,false , false, false, false, false},
				},
				row: 6,
				col: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfGoodChoice(tt.args.grid, tt.args.col, tt.args.row); got != tt.want {
				t.Errorf("checkIfGoodChoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
