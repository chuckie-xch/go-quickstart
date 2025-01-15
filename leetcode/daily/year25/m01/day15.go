package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func minOperations(nums []int, k int) (ans int) {
	pq := &minHeap{nums}
	heap.Init(pq)
	for ; pq.Len() > 1 && pq.IntSlice[0] < k; ans++ {
		x, y := heap.Pop(pq).(int), heap.Pop(pq).(int)
		heap.Push(pq, min(x, y)*2+max(x, y))
	}
	return
}

type minHeap struct {
	sort.IntSlice
}

func (h *minHeap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func (h *minHeap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

func (h *minHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
