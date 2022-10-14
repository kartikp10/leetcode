package leetcode

/*
 * Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
 * A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    var helper func(int, int) *TreeNode
    helper = func(left int, right int) *TreeNode {
        if left > right {
            return nil
        }
        p := (left + right)/2
        root := &TreeNode{
        	Val:   nums[p],
        	Left:  nil,
        	Right: nil,
        }
        root.Left = helper(left, p-1)
        root.Right = helper(p+1, right)
        return root
    }
    return helper(0, len(nums)-1)
}
