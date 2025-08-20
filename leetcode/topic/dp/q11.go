package dp

func robII(nums []int) int {
	n := len(nums)
	return max(nums[0]+rob1(nums, 2, n-2), rob1(nums, 1, n-1))
}

func rob1(nums []int, start int, end int) int {
	f0, f1 := 0, 0
	for i := start; i <= end; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}
