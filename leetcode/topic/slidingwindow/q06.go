package main

func maxScore(cardPoints []int, k int) int {
	s := 0
	n := len(cardPoints)
	for i := 0; i < k; i++ {
		s += cardPoints[i]
	}
	ans := s

	for i := 1; i <= k; i++ {
		s += cardPoints[n-i] - cardPoints[k-i]
		ans = max(ans, s)
	}

	return ans
}
