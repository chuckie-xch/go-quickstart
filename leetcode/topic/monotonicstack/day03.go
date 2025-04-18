package monotonicstack

func carFleet(target int, position []int, speed []int) int {
	t := make([]float32, target)
	st := []int{}
	for i, v := range position {
		t[position[i]] = float32(target-v) / float32(speed[i])
	}
	for i, v := range t {
		if v != 0 {
			for len(st) > 0 && v >= t[st[len(st)-1]] {
				st = st[:len(st)-1]
			}
			st = append(st, i)
		}
	}
	return len(st)
}
