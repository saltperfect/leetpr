package tree

import "fmt"

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func Init() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(a ...int) {
	for _, val := range a {
		insert(b.Root, val)
	}
}

func insert(root *TreeNode, target int) {
	if root == nil {
		root = &TreeNode{Val: target}
		return
	}
	if root.Val > target {
		if root.Left == nil {
			root.Left = &TreeNode{Val: target}
		}
		insert(root.Left, target)
	}
	if root.Val < target {
		if root.Right == nil {
			root.Right = &TreeNode{Val: target}
		}
		insert(root.Right, target)
	}
}

func (b *BinaryTree) Print() {
	printInorder(b.Root)
} 

func printInorder(root *TreeNode) {
	if root  == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d, ", root.Val)
	printInorder(root.Right)	
}

func (b *BinaryTree) Delete(a ...int) {
	for _, val := range a {
		deleteRoot(b.Root, val)
	}
}

func deleteRoot(root *TreeNode, a int) {

}