package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

func ListToNum(l *ListNode) int {
	num := 0.00
	pow := 0
	for l != nil {
		num = num + float64(l.Val)*math.Pow10(pow)
		pow++
		l = l.Next
	}
	return int(num)
}

func StrToList(sn string) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: &ListNode{},
	}
	curr := head
	for i := len(sn) - 1; i >= 0; i-- {
		curr.Val, _ = strconv.Atoi(string(sn[i]))
		if i == 0 {
			curr.Next = nil
			continue
		}
		next := &ListNode{
			Val:  0,
			Next: nil,
		}
		curr.Next = next
		curr = next
	}
	return head
}

func ArrToList(arr []int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: &ListNode{},
	}
	curr := head
	for i, v := range arr {
		curr.Val = v
		if i == len(arr)-1 {
			curr.Next = nil
			continue
		}
		next := &ListNode{
			Val:  0,
			Next: nil,
		}
		curr.Next = next
		curr = next
	}
	return head
}

func PrintList(l *ListNode) {
	for l.Next != nil {
		fmt.Println(l.Val, " -> ")
		l = l.Next
	}
	if l.Next == nil {
		fmt.Println(l.Val, " nil")
		return
	}

}
