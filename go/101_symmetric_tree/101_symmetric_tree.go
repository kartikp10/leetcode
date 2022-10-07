package leeetcode

// Given the root of a binary tree, check whether it is a mirror of
// itself (i.e., symmetric around its center).
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && isMirror(l.Right, r.Left) && isMirror(l.Left, r.Right)
}

//                     1
//                 /       \
//                2         2
//               /   \     /   \
//              3    4     4    3
//             / \  / \   / \  / \
//            5  6 7  8  8  7 6  5
