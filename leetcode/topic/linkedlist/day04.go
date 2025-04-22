package linkedlist

import "go-quickstart/leetcode/datastruct"

func middleNode(head *datastruct.ListNode) *datastruct.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
