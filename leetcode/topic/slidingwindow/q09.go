package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	ans, l := 0, 0
	cnt := make([]int, 128)
	cnt[s[l]]++
	for r := 1; r < n; r++ {
		cur := s[r]
		cnt[cur]++
		for cnt[cur] > 1 {
			cnt[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func lengthOfLongestSubstring2(s string) (ans int) {
	l := 0
	cnt := [128]int{}
	for r, cur := range s {
		cnt[cur]++
		for cnt[cur] > 1 {
			cnt[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}
