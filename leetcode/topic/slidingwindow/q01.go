package main

func main() {

}

func maxVowels(s string, k int) int {
	ans, vowel := 0, 0
	vowelMap := make(map[rune]bool)
	for _, v := range []rune{'a', 'e', 'i', 'o', 'u'} {
		vowelMap[v] = true
	}

	for i := 0; i < len(s); i++ {
		if vowelMap[rune(s[i])] {
			vowel++
		}
		if i+1 < k {
			continue
		}
		ans = max(ans, vowel)
		out := s[i-k+1]
		if vowelMap[rune(out)] {
			vowel--
		}
	}
	return ans
}
