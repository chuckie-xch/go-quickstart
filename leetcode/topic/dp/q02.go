package dp

func minCostClimbingStairs(cost []int) int {
	f0, f1, n := 0, 0, len(cost)
	for i := 2; i <= n; i++ {
		newF := min(f1+cost[i-1], f0+cost[i-2])
		f0 = f1
		f1 = newF
	}
	return f1
}
