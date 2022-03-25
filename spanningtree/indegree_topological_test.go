package spanningtree

import "testing"

func Test_tropicalsort(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				grid: [][]int{{1,4}, {2}, {3},{},{3},{1,2,4}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tropicalsort(tt.args.grid); got != tt.want {
				t.Errorf("tropicalsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
