package monotonicstack

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range nums {
		ans[i] = -1
	}
	st := []int{}
	for i := 0; i < 2*n; i++ {
		index := i % n
		x := nums[index]
		for len(st) > 0 && x > nums[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = x
		}
		if i < n {
			st = append(st, i)
		}
	}
	return ans

}
