package main

func main() {

}

func minimumSubarrayLength2(nums []int, k int) int {
	n := len(nums)
	l, r := 0, -1
	ans := n + 1
	bit := make([]int, 32)

	for r < n {
		for check(bit, k) && l <= r {
			ans = min(ans, r-l+1)
			for i := 0; i <= 30; i++ {
				bit[i] -= (nums[l] >> i) & 1
			}
			l++
		}
		r++
		if r == n {
			break
		}
		for i := 0; i <= 30; i++ {
			bit[i] += (nums[r] >> i) & 1
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}

func check(bit []int, k int) bool {
	for i := 30; i >= 0; i-- {
		if bit[i] > 0 && ((k>>i)&1) == 0 {
			return true
		}
		if bit[i] == 0 && ((k>>i)&1) == 1 {
			return false
		}
	}
	return true
}
