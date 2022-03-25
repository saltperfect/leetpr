package spanningtree

import "testing"

func Test_prim(t *testing.T) {
	type args struct {
		grid [][][2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				grid: [][][2]int{
					{{1,1},{6,2},{5,3}},
					{{2,2},{0,1},{5,1}},
					{{1,2},{3,6}},
					{{5,3},{2,6},{4,3}},
					{{3,3},{6,1},{5,4}},
					{{0,3},{1,1},{3,3}, {4,4}, {6,1}},
					{{0,2},{5,1},{4,1}},
				},
			},
			want: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prim(tt.args.grid); got != tt.want {
				t.Errorf("prim() = %v, want %v", got, tt.want)
			}
		})
	}
}
