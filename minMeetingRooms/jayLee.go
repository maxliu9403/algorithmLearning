package minMeetingRooms

import (
	"container/heap"
	"sort"
)

/*
Question
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],…] (si < ei)，为避免会议冲突，同时要考虑充分利用会议室资源，
请你计算至少需要多少间会议室，才能满足这些会议安排。

Example 1:
Input: [[0, 30],[5, 10],[15, 20]]
Output: 2

Example 2:
Input: [[7,10],[2,4]]
Output: 1

Example 3:
Input: [[1,10],[2,7],[3,19],[8,12],[10,20],[11,30]]
Output: 4

Thinking:
1. 首先,确保所有会议的开始时间是升序的;
2. 其次,需要选取其中最先结束的会议,并寻找比它晚开始的会议,把其也安排在这间会议室,这样就可以节省一间.
3. 另外,需要注意的是,会议室被安排后,需要更新该会议室的结束时间,并重新选取最先结束的会议,重复第2步...
*/
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &IntHeap{} // 实现堆操作的5个方法(必须指针类型)
	heap.Init(h)    // 初始化堆(必须有这一步)
	for _, interval := range intervals {
		if len(*h) == 0 {
			heap.Push(h, interval[1])
			continue
		}
		// 当前会议的开始时间正好比最早结束的会议晚,可以共用一间会议室
		if interval[0] >= (*h)[0] {
			_ = heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}
	return len(*h)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
