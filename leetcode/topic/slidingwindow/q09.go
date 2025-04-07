package main

func maximumLengthSubstring(s string) int {
	ans, l := 0, 0
	cnt := make([]int, 128)
	for r, cur := range s {
		cnt[cur]++
		for cnt[cur] > 2 {
			cnt[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
