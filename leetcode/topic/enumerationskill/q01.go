package enumerationskill

func maxScoreSightseeingPair(values []int) int {
	ans, maxLeft := 0, 0
	for i, v := range values {
		ans = max(ans, v-i+maxLeft)
		maxLeft = max(maxLeft, v+i)
	}
	return ans
}
