package enumerationskill

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	cnt := map[int]int{}
	for _, a := range nums1 {
		for _, b := range nums2 {
			cnt[a+b]++
		}
	}
	ans := 0
	for _, a := range nums3 {
		for _, b := range nums4 {
			ans += cnt[-a-b]
		}
	}
	return ans
}
