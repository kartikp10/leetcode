package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// slow and fast pointer approach
// tortoise-and-hare algorithm
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	pre, slow, fast := head, head, head
	for slow.Next != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next

		if fast.Next != nil {
			fast = fast.Next
		}
	}
	pre.Next = slow.Next
	return head
}

func deleteMiddle2(head *ListNode) *ListNode {
	length := getLen(head)
	if length == 1 {
		return nil
	} else if length == 2 {
		head.Next = nil
	} else {
		deleteNode(getNodeNum(length/2, head))
	}
	return head
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

func getLen(l *ListNode) int {
	length := 1
	for l.Next != nil {
		length++
		l = l.Next
	}
	return length
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
