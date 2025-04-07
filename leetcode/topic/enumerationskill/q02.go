package enumerationskill

import "math"

func minimumCardPickup(cards []int) int {
	ans := math.MaxInt
	positionMap := map[int]int{}
	for i, v := range cards {
		if pos, ok := positionMap[v]; ok {
			ans = min(ans, i-pos+1)
		}
		positionMap[v] = i
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
