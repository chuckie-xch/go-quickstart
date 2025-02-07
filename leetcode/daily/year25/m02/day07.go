package main

func main() {
}

func generateMatrix(n int) [][]int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	i, j, di := 0, 0, 0
	for val := 1; val <= n*n; val++ {
		ans[i][j] = val
		x := i + dirs[di][0]
		y := j + dirs[di][1]
		if x < 0 || x >= n || y < 0 || y >= n || ans[x][y] != 0 {
			di = (di + 1) % 4
		}
		i += dirs[di][0]
		j += dirs[di][1]
	}
	return ans
}
