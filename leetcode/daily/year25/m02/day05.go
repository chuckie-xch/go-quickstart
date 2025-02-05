package main

import "slices"

func main() {

}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	length := len(nums)
	path := []int{}
	ans := [][]int{}
	var dfs func(int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, slices.Clone(path))
			return
		}
		x := nums[i]
		path = append(path, x)
		dfs(i + 1)
		path = path[:len(path)-1]

		i++
		for i < length && nums[i] == x {
			i++
		}
		dfs(i)

	}
	dfs(0)
	return ans
}
