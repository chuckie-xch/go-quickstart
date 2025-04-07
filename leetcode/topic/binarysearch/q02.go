package binarysearch

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	l, r := 0, n-1
	for l <= r && r-l+1 > k {
		if abs(arr[l]-x) <= abs(arr[r]-x) {
			r--
		} else {
			l++
		}
	}
	return arr[l : r+1]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
