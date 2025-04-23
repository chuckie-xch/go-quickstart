package linkedlist

import "go-quickstart/leetcode/datastruct"

func removeNthFromEnd(head *datastruct.ListNode, n int) *datastruct.ListNode {
	dummy := &datastruct.ListNode{Next: head}
	left, right := dummy, dummy
	for n > 0 {
		right = right.Next
		n--
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
