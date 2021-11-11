package dfs

import (
	"fmt"
	"testing"
)

func Test_dfst(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		args args
	}{
		{
			args: args{
				grid: [][]int{
					{1,7},
					{2},
					{5},
					{1,4},
					{5},
					{},
					{7},
					{},
					{},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dfst(tt.args.grid)
		})
	}
}
