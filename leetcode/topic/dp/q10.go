package dp

func maxProduct(nums []int) int {
	n := len(nums)
	preMax, preMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		curMax := max(x, x*preMax, x*preMin)
		curMin := min(x, x*preMax, x*preMin)
		ans = max(ans, curMax)
		preMax = curMax
		preMin = curMin
	}
	return ans
}
