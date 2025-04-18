package monotonicstack

import "slices"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	st := []int{}
	for i, x := range slices.Backward(temperatures) {
		for len(st) > 0 && x >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return ans
}

func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	st := []int{}
	for i, x := range temperatures {
		for len(st) > 0 && x > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}
