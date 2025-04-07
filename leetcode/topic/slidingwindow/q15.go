package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	start, end, need := -1, len(s), len(t)
	valueMap := make([]int, 128)
	for _, c := range t {
		valueMap[c]++
	}

	l := 0
	for r, k := range s {
		valueMap[k]--
		if valueMap[k] >= 0 {
			need--
		}
		if need == 0 {
			for valueMap[s[l]] < 0 {
				valueMap[s[l]]++
				l++
			}
			if end-start > r-l {
				start, end = l, r
			}
			valueMap[s[l]]++
			l++
			need++
		}
	}

	if start < 0 {
		return ""
	}

	return s[start : end+1]
}
