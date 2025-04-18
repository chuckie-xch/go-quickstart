package m04

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n%2 != 0 {
			continue
		}
		diff := 0
		for _, x := range s[:n/2] {
			diff += int(x)
		}
		for _, x := range s[n/2:] {
			diff -= int(x)
		}
		if diff == 0 {
			ans++
		}
	}
	return ans
}
