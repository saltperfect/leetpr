package binarysearch

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getSoldiers(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 0, 0, 0},
			want: 2,
		},
		{
			nums: []int{1, 1, 1, 1, 0},
			want: 4,
		},
		{
			nums: []int{1, 1, 1, 1, 1},
			want: 5,
		},
		{
			nums: []int{0, 0, 0, 0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			if got := getSoldiers(tt.nums); got != tt.want {
				t.Errorf("getSoldiers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		num [][2]int
		k   [2]int
	}
	tests := []struct {
		args args
		want [][2]int
	}{
		{
			args: args{
				num: [][2]int{{4, 1}, {6, 2}, {9, 3}},
				k : [2]int {7, 4},
			},
			want: [][2]int{{4, 1}, {6, 2}, {7, 4}, {9, 3} },
		},
		{
			args: args{
				num: [][2]int{{4, 1}, {6, 2}, {9, 3}},
				k : [2]int {1, 4},
			},
			want: [][2]int{{1, 4}, {4, 1}, {6, 2}, {9, 3} },
		},
		{
			args: args{
				num: [][2]int{{4, 1}, {6, 2}, {9, 3}},
				k : [2]int {6, 4},
			},
			want: [][2]int{{4, 1}, {6, 2}, {6, 4}, {9, 3} },
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := insert(tt.args.num, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
