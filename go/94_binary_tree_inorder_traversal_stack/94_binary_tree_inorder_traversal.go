package leetcode

import "fmt"

type Stack []int

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		i := len(*s) - 1
		e := (*s)[i]
		*s = (*s)[:i]
		return e, true
	}
}

func inorderTraversal(root *TreeNode) []int {
	// add current and go to left until there is no left
	// pop
	// go to right until there is no right
	// var io []int
	/* TODO: */
	stk := Stack{}
	var traverseLeft func(*TreeNode, Stack)

	traverseLeft = func(tn *TreeNode, s Stack) {
		if tn != nil {
			s.Push(tn.Val)
			tn = tn.Left
			traverseLeft(tn, s)
			if v, ok := s.Pop(); ok {
				fmt.Printf("Popping: %v \n", v)
			}
		}
	}
	if root == nil {
		return make([]int, 0)
	} else {
		traverseLeft(root, stk)
	}
	return make([]int, 0)
}
