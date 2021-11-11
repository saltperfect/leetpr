package sorting

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "even",
			args: args{
				nums: []int{8, 12, 2, 5, 18, 13, 6, 4},
			},
			want: []int{2, 4, 5, 6, 8, 12, 13, 18},
		},
		{
			name: "odd",
			args: args{
				nums: []int{8, 12, 2, 5, 18, 13, 6},
			},
			want: []int{2, 5, 6, 8, 12, 13, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertionsort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "even",
			args: args{
				nums: []int{8, 12, 2, 5, 18, 13, 6, 4},
			},
			want: []int{2, 4, 5, 6, 8, 12, 13, 18},
		},
		{
			name: "odd",
			args: args{
				nums: []int{8, 12, 2, 5, 18, 13, 6},
			},
			want: []int{2, 5, 6, 8, 12, 13, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionsort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
