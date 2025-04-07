package main

func minOperations(nums []int, x int) int {
	target := -x
	for _, num := range nums {
		target += num
	}
	if target < 0 {
		return -1
	}
	l, s, ans := 0, 0, -1
	for r, cur := range nums {
		s += cur
		for s > target {
			s -= nums[l]
			l++
		}
		if s == target {
			ans = max(ans, r-l+1)
		}
	}
	if ans == -1 {
		return ans
	}
	return len(nums) - ans
}
