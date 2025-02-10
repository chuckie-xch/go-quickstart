package main

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || j < 0 || obstacleGrid[i][j] == 1 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		memo[i][j] = dfs(i, j-1) + dfs(i-1, j)
		return memo[i][j]
	}
	return dfs(m-1, n-1)
}
