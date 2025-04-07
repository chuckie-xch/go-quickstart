package doublepointer

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	m, n := len(a), len(b)
	if m == 1 && n == 1 {
		return int(math.Abs(float64(a[0] - b[0])))
	}

	i, j, ans := 0, 0, math.MaxInt64
	sort.Ints(a)
	sort.Ints(b)
	for i < m && j < n {
		if a[i] == b[j] {
			return 0
		}
		ans = min(ans, int(math.Abs(float64(a[i]-b[j]))))
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func smallestDifference2(a []int, b []int) int {
	m, n := len(a), len(b)
	if m == 1 && n == 1 {
		return abs(a[0] - b[0])
	}

	i, j, ans := 0, 0, math.MaxInt64
	sort.Ints(a)
	sort.Ints(b)
	for i < m && j < n {
		if a[i] == b[j] {
			return 0
		}
		ans = min(ans, abs(a[i]-b[j]))
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
