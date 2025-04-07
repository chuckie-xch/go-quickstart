package m03

func sumOfBeauties(nums []int) int {
	n := len(nums)
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	ans, preMax := 0, nums[0]
	for i := 1; i < n-1; i++ {
		if preMax < nums[i] && nums[i] < sufMin[i+1] {
			ans += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans += 1
		}
		preMax = max(preMax, nums[i])
	}
	return ans
}
