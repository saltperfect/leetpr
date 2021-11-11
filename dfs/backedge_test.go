package dfs

import (
	"fmt"
	"testing"
)

func Test_dfs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		args args
	}{
		{
			args: args{
				grid: [][]int{
					{1,3},
					{4},
					{4,5},
					{1},
					{3},
					{5},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dfs(tt.args.grid)
		})
	}
}
