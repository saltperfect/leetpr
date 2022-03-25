package sorting

import (
	"reflect"
	"testing"
)

func Test_mymerge(t *testing.T) {
	type args struct {
		nums  *[]int
		left  int
		right int
		mid   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums:  &[]int{0, 0, 2, 4, 5, 6, 1, 3, 7},
				left:  2,
				right: 8,
				mid:   5,
			},
		},
		{
			name: "second",
			args: args{
				nums:  &[]int{2, 3, 4, 5, 1},
				left:  3,
				right: 4,
				mid:   3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mymerge(tt.args.nums, tt.args.left, tt.args.mid, tt.args.right)
		})
	}
}


func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{5,2,6,1},
			},
			want: []int{2,1,1,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
