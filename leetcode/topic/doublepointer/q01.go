package doublepointer

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	for i := n - 1; i >= 2; i-- {
		k := nums[i]
		l, r := 0, i-1

		for l < r {
			s := nums[l] + nums[r]
			if s <= k {
				l++
			} else {
				ans += r - l
				r--
			}
		}
	}

	return ans
}
