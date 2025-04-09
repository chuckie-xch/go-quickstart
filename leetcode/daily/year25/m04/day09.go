package m04

func minOperations(nums []int, k int) int {
	st := map[int]struct{}{}
	for _, v := range nums {
		if v < k {
			return -1
		}
		if v > k {
			st[v] = struct{}{}
		}
	}
	return len(st)
}
