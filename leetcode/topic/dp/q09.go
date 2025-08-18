package dp

import "slices"

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int64, n+1)
	for i, q := range slices.Backward(questions) {
		j := min(n, q[1]+i+1)
		f[i] = max(f[i+1], int64(q[0])+f[j])
	}
	return f[0]
}
