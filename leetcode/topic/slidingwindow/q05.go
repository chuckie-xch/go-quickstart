package main

func main() {

}

func maximumSubarraySum(nums []int, k int) int64 {
	ans, s := 0, 0
	cnt := map[int]int{}
	for _, x := range nums[:k-1] {
		cnt[x]++
		s += x
	}
	for i := k - 1; i < len(nums); i++ {
		cnt[nums[i]]++
		s += nums[i]
		if len(cnt) == k {
			ans = max(ans, s)
		}
		x := nums[i-k+1]
		s -= x
		cnt[x]--
		if cnt[x] == 0 {
			delete(cnt, x)
		}
	}
	return int64(ans)
}
