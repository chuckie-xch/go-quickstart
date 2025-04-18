package linkedlist

import "go-quickstart/leetcode/datastruct"

func reverseList(head *datastruct.ListNode) *datastruct.ListNode {
	var pre, cur *datastruct.ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
