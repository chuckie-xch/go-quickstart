package main

func main() {

}

func findSpecialInteger(arr []int) int {
	n := len(arr)
	lower := n / 4
	total := 1
	pre := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] == pre {
			total++
		} else {
			total = 1
			pre = arr[i]
		}
		if total > lower {
			break
		}
	}
	return pre
}
