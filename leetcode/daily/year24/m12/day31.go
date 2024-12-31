package main

import "slices"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) (ans int64) {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	i, j := 0, 0
	for range m + n - 2 {
		if j == n-1 || i < m-1 && horizontalCut[i] < verticalCut[j] {
			ans += int64(horizontalCut[i] * (n - j))
			i++
		} else {
			ans += int64(verticalCut[j] * (m - i))
			j++
		}
	}
	return
}

func main() {

}
