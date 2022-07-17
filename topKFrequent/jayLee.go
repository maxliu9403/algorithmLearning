package topKFrequent

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	h := map[int]int{}
	for _, num := range nums {
		h[num]++
	}
	hp := &hp{}
	heap.Init(hp)
	for num, count := range h {
		heap.Push(hp, [2]int{num, count})
		if hp.Len() > k {
			heap.Pop(hp) // 次数最小的PASS
		}
	}
	var resp []int
	for hp.Len() > 0 {
		resp = append(resp, heap.Pop(hp).([2]int)[0])
	}
	return resp
}

type hp [][2]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *hp) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
