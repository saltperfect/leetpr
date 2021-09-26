package main

import (
	"fmt"
	"testing"
)

func Test_smallestDivisor(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1,2,5,9},
				threshold: 6,
			},
			want: 5,			
		},
		{
			args: args{
				nums: []int{21212,10101,12121},
				threshold: 1000000,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{2,3,5,7,11},
				threshold: 11,
			},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i) , func(t *testing.T) {
			if got := smallestDivisor(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("smallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
