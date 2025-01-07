package main

import "strings"

func countKeyChanges(s string) int {
	ans := 0
	s = strings.ToLower(s)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return ans
}
func main() {

}
