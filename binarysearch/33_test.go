package binarysearch

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{4,5,6,7,0,1,2},
				target: 0,
			},
			want: 4,

		},
		{
			args: args{
				nums: []int{4,5,6,7,0,1,2},
				target: 3,
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
