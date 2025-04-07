package enumerationskill

func maxDistance(arrays [][]int) int {
	ans := 0
	preMin := arrays[0][0]
	preMax := arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		a := arrays[i]
		x, y := a[0], a[len(a)-1]
		ans = max(ans, preMax-x, y-preMin)
		preMax = max(preMax, y)
		preMin = min(preMin, x)
	}
	return ans
}
