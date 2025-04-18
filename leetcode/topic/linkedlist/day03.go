package linkedlist

import "go-quickstart/leetcode/datastruct"

func reverseBetween(head *datastruct.ListNode, left int, right int) *datastruct.ListNode {
	dummy := &datastruct.ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	var pre *datastruct.ListNode = nil
	cur := p0.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}
