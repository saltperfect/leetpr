package tree

import "testing"

func TestBinaryTree_Insert(t *testing.T) {
	type fields struct {
		Root *TreeNode
	}
	type args struct {
		a []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "first",
			args: args{
				a: []int{5, 7, 1, 34, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{Root: &TreeNode{Val: 2}}
			b.Insert(tt.args.a...)
			b.Print()
		})
	}
}
