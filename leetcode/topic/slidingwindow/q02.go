package main

import "math"

func main() {

}

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i+1 < k {
			continue
		}
		maxSum = max(maxSum, sum)
		sum -= nums[i-k+1]
	}
	return float64(maxSum) / float64(k)
}
