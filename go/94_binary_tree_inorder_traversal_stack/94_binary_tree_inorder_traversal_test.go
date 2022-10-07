package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTree(data []int) *TreeNode {
	var root *TreeNode
	var curr *TreeNode
	curr = root
	diff := 0
	for i, d := range data {
		curr.Val = d
		curr.Left = &TreeNode{Val: data[i+int(diff)], Left: nil, Right: nil}
		curr.Right = &TreeNode{Val: data[i+int(diff)+1], Left: nil, Right: nil}
        diff++
	}
	return root
}
