package main

func main() {

}

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	ans := n + 1
	for i := 0; i < n; i++ {
		value := 0
		for j := i; j < n; j++ {
			value = value | nums[j]
			if value >= k {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}
