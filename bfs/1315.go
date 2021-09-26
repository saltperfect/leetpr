package bfs

import (
	"github.com/saltperfect/goleet"
)

func findSum(root *goleet.TreeNode) int {
	sum := 0
	if root.Left != nil {
		if root.Left.Left != nil {
			sum += root.Left.Left.Val
		}
		if root.Left.Right != nil {
			sum += root.Left.Right.Val
		}
	}
	if root.Right != nil {
		if root.Right.Left != nil {
			sum += root.Right.Left.Val
		}
		if root.Right.Right != nil {
			sum += root.Right.Right.Val
		}
	}
	return sum
}

func sumEvenGrandparent(root *goleet.TreeNode) int {
	var q, nq []*goleet.TreeNode
	nq = []*goleet.TreeNode{root}
	sum := 0
	for len(nq) > 0 {
		q, nq = nq, make([]*goleet.TreeNode, 0)
		for _, v := range q {
			if v.Val%2 == 0 {
				sum += findSum(v)
			}
			if v.Left != nil {
				nq = append(nq, v.Left)
			}
			if v.Right != nil {
				nq = append(nq, v.Right)
			}
		}
	}
	return sum
}
