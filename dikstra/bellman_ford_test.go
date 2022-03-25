package dikstra

import "testing"

func Test_bellman(t *testing.T) {
	type args struct {
		grid [][][2]int
		root int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				grid: [][][2]int{
					{ {1,1}, {4,2} },
					{ {2,3} },
					{ {3,2} },
					{ {5,2} },
					{ {1,-2}, {5,3} },
					{ {1,-6}, {2,-4} },
				},
				root: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bellman(tt.args.grid, tt.args.root)
		})
	}
}
