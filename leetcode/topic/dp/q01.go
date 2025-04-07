package dp

import "math"

func palindromePartition(s string, k int) int {
	n := len(s)
	memoChange := make([][]int, n)
	for i := range memoChange {
		memoChange[i] = make([]int, n)
		for j := range memoChange[i] {
			memoChange[i][j] = -1
		}
	}

	var minChange func(int, int) int
	minChange = func(i int, j int) int {
		if memoChange[i][j] != -1 {
			return memoChange[i][j]
		}
		if i >= j {
			return 0
		}
		res := minChange(i+1, j-1)
		if s[i] != s[j] {
			res++
		}
		memoChange[i][j] = res
		return res
	}

	memoDfs := make([][]int, n)
	for i := range memoDfs {
		memoDfs[i] = make([]int, n)
		for j := range memoDfs[i] {
			memoDfs[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i int, r int) int {
		if memoDfs[i][r] != -1 {
			return memoDfs[i][r]
		}
		if i == 0 {
			return minChange(0, r)
		}
		res := math.MaxInt
		for l := i; l <= r; l++ {
			res = min(res, minChange(l, r)+dfs(i-1, l-1))
		}
		memoDfs[i][r] = res
		return res
	}
	return dfs(k-1, n-1)

}
