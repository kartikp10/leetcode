package leetcode

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

// create list from int slice and return head
func createList(nums []int) *ListNode {
    head := new(ListNode)
    dummy := head
    for i, n := range nums {
        dummy.Val = n
        if i == len(nums)-1 {
            dummy.Next = nil
        } else {
            dummy.Next = &ListNode{
                Val:  0,
                Next: &ListNode{},
            }
            dummy = dummy.Next
        }
    }
    return head
}

func listToSlice(l *ListNode) []int {
    var nums []int
	for l.Next != nil {
        nums = append(nums, l.Val)
		l = l.Next
	}
    return nums
}

func printList(l *ListNode) {
    nums := listToSlice(l)
    for i, n := range nums {
        fmt.Print(n, " -> ")
        if i == len(nums)-1 {
            fmt.Print("nil\n")
        }
    }
}

func getNodeNum(n int, head *ListNode) *ListNode {
	dummy := head
	for i := 0; i <= n-1; i++ {
		if i == n-1 {
			return dummy
		}
		dummy = dummy.Next
	}
	return nil
}
