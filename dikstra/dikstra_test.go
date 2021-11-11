package dikstra

import (
	"fmt"
	"testing"
)

func Test_dikstra(t *testing.T) {
	type args struct {
		graph [][]vertex
		start int
	}
	tests := []struct {
		args args
	}{
		{
			args: args{
				graph: [][]vertex{
					{{1,19},{2,7}},
					{{0,19},{3,4},{2,11}},
					{{0,7},{1,11},{3,15},{4,5}},
					{{1,4},{2,15},{4,13}},
					{{2,5},{3,13}},
				},
				start: 0,
			},
		},
		{
			args: args{
				graph: [][]vertex{
					{{1,19},{2,7}},
					{{0,19},{3,4},{2,11}},
					{{0,7},{1,11},{3,15},{4,5}},
					{{1,4},{2,15},{4,13}},
					{{2,5},{3,13}},
				},
				start: 3,
			},
		},
		{
			args: args{
				graph: [][]vertex{
					{{1,19},{2,7}},
					{{0,19},{3,4},{2,11}},
					{{0,7},{1,11},{3,15},{4,5}},
					{{1,4},{2,15},{4,13}},
					{{2,5},{3,13}},
				},
				start: 4,
			},
		},
		// TODO: Add test cases.
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dikstra(tt.args.graph, tt.args.start)
		})
	}
}
