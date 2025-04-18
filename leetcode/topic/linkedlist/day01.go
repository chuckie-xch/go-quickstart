package linkedlist

import "go-quickstart/leetcode/datastruct"

func getDecimalValue(head *datastruct.ListNode) int {
	ans := 0
	temp := head
	for temp != nil {
		ans = ans*2 + temp.Val
		temp = temp.Next
	}
	return ans
}
