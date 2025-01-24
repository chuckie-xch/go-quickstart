package main

func main() {

}

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
	}
	return dfs(n-1, k, piles, memo)
}

func dfs(i int, j int, piles [][]int, memo [][]int) int {
	if i < 0 {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	ans := dfs(i-1, j, piles, memo)

	totalValue := 0
	for w := 0; w < min(j, len(piles[i])); w++ {
		totalValue += piles[i][w]
		ans = max(ans, dfs(i-1, j-w-1, piles, memo)+totalValue)
	}

	memo[i][j] = ans
	return ans
}
