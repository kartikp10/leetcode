package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	io := []int{}
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root != nil {
			helper(root.Left)
			io = append(io, root.Val)
			helper(root.Right)
		}
	}
	helper(root)
	return io
}
