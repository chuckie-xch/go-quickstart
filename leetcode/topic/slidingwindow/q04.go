package main

func main() {

}

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	n := len(nums)
	s := 0
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < n; i++ {
		s += nums[i]
		if i < 2*k {
			continue
		}
		ans[i-k] = s / (2*k + 1)
		s -= nums[i-2*k]
	}
	return ans
}
