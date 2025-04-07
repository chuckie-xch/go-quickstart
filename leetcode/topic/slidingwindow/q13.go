package main

import (
	"slices"
)

func treeSum(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i < n-2; i++ {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+1] > 0 {
			break
		}
		if x+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			s := x + nums[l] + nums[r]
			if s > 0 {
				r--
			} else if s < 0 {
				l++
			} else {
				ans = append(ans, []int{x, nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ {

				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {

				}
			}
		}
	}

	return ans

}
