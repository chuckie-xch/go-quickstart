package main

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	var pos []int
	for i, num := range nums {
		if num == x {
			pos = append(pos, i)
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q > len(pos) || q < 1 {
			ans[i] = -1
		} else {
			ans[i] = pos[q-1]
		}
	}
	return ans
}

func main() {
	// 创建一些测试案例
	testCases := []struct {
		nums    []int
		queries []int
		x       int
		want    []int
	}{
		{[]int{1, 2, 3, 4, 5, 2, 3}, []int{1, 2, 5}, 2, []int{1, 5, -1}},
		{[]int{1, 3, 5, 7}, []int{1, 2}, 2, []int{-1, -1}},
		{[]int{8, 8, 8, 8, 8}, []int{1, 3, 5, 6}, 8, []int{0, 2, 4, -1}},
		{[]int{1, 2, 2, 2, 3}, []int{1, 2, 3, 4}, 2, []int{1, 2, 3, -1}},
		{[]int{1, 2, 3, 4, 5}, []int{0, -1, 6}, 3, []int{-1, -1, -1}},
	}

	// 执行测试案例
	for _, tc := range testCases {
		got := occurrencesOfElement(tc.nums, tc.queries, tc.x)
		fmt.Printf("For nums=%v, queries=%v, and x=%d, got %v, want %v\n", tc.nums, tc.queries, tc.x, got, tc.want)
	}
}
