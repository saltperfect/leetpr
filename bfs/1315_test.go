package bfs

import (
	"testing"

	"github.com/saltperfect/goleet"
)

func Test_findSum(t *testing.T) {
	type args struct {
		root *goleet.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenGrandparent(tt.args.root); got != tt.want {
				t.Errorf("findSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
