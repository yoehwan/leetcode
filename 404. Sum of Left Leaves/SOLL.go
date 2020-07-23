package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	fmt.Println(sumOfLeftLeaves(tree))

}

// Definition for a binary tree node.

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}
	if root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)

}
