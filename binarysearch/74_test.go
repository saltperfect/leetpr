package binarysearch

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
