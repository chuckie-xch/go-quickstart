package prefixsum

func maxAbsoluteSum(nums []int) int {
	s, mx, mn := 0, 0, 0
	for _, x := range nums {
		s += x
		mx = max(mx, s)
		mn = min(mn, s)
	}
	return mx - mn
}
