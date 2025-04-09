package prefixsum

import "math"

func maxSubArray(nums []int) int {
	ans, mn, s := math.MinInt, 0, 0
	for _, x := range nums {
		s += x
		ans = max(ans, s-mn)
		mn = min(mn, s)
	}
	return ans
}
