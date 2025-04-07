package main

func trap(height []int) (ans int) {
	n := len(height)
	if n < 3 {
		return
	}
	lMax, rMax, l, r := height[0], height[n-1], 0, n-1
	for l <= r {
		if lMax < rMax {
			ans += max(0, lMax-height[l])
			lMax = max(lMax, height[l])
			l++
		} else {
			ans += max(0, rMax-height[r])
			rMax = max(rMax, height[r])
			r--
		}
	}

	return
}
