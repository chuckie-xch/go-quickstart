package main

func main() {

}

func numOfSubArrays(arr []int, k int, threshold int) int {
	ans, sum := 0, 0
	threshold = threshold * k
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i+1 < k {
			continue
		}
		if sum >= threshold {
			ans++
		}
		sum -= arr[i-k+1]
	}
	return ans
}
