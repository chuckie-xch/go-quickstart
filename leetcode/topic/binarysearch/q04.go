package binarysearch

import (
	"slices"
)

func maximumTastiness(price []int, k int) int {
	slices.Sort(price)
	var f func(int) int
	f = func(d int) int {
		cnt, pre := 1, price[0]
		for _, p := range price {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt
	}
	l, r := 0, (price[len(price)-1]-price[0])/(k-1)
	for l <= r {
		mid := l + ((r - l) >> 1)
		if f(mid) >= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1

}
