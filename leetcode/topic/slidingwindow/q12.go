package main

func minimumRefill(plants []int, capacityA int, capacityB int) (ans int) {
	l, r, totalA, totalB := 0, len(plants)-1, capacityA, capacityB
	for l < r {
		if totalA < plants[l] {
			ans++
			totalA = capacityA
		}
		if totalB < plants[r] {
			ans++
			totalB = capacityB
		}
		totalA -= plants[l]
		totalB -= plants[r]
		l++
		r--
	}

	if l == r && max(totalA, totalB) < plants[l] {
		ans++
	}

	return
}
