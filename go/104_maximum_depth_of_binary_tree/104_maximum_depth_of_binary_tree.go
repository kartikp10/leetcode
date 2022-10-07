package leetcode

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the
// longest path from the root node down to the farthest leaf node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left_depth := maxDepth(root.Left)
		right_depth := maxDepth(root.Right)
		return max(left_depth, right_depth) + 1
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
