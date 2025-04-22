package linkedlist

import "go-quickstart/leetcode/datastruct"

func isPalindrome(head *datastruct.ListNode) bool {
	mid := findMid(head)
	head2 := getReverseList(mid)
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func findMid(head *datastruct.ListNode) *datastruct.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func getReverseList(head *datastruct.ListNode) *datastruct.ListNode {
	var pre, next *datastruct.ListNode = nil, nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
