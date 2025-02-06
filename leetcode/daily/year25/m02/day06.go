package main

import "slices"

func main() {

}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	ans := [][]int{}
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j, on := range onPath {
			if on || (j > 0 && nums[j] == nums[j-1] && !onPath[j-1]) {
				continue
			}
			onPath[j] = true
			path[i] = nums[j]
			dfs(i + 1)
			onPath[j] = false
		}
	}
	dfs(0)
	return ans
}
