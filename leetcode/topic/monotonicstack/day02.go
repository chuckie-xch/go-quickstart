package monotonicstack

import "slices"

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{}
	for i, x := range prices {
		for len(st) > 0 && x <= prices[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = prices[j] - x
		}
		st = append(st, i)
	}
	for len(st) > 0 {
		i := st[len(st)-1]
		st = st[:len(st)-1]
		ans[i] = prices[i]
	}
	return ans
}

func finalPrices2(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{}
	for i, x := range slices.Backward(prices) {
		for len(st) > 0 && x < prices[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = x - prices[st[len(st)-1]]
		} else {
			ans[i] = prices[i]
		}
		st = append(st, i)
	}
	return ans
}
