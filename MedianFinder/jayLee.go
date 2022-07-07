package MedianFinder

import (
	"container/heap"
	"sort"
)

// MedianFinder
/**
- void addNum(int num) - 从数据流中添加一个整数到数据结构中。
- double findMedian() - 返回目前所有元素的中位数
*/
type MedianFinder struct {
	A minHp // A: 最小堆，记录大于中位数的所有元素
	B maxHp // B: 最大堆，记录小于等于中位数的所有元素，父节点为中位数
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

// AddNum 判断num放入A还是放入B
func (this *MedianFinder) AddNum(num int) {
	maxQ, minQ := &this.A, &this.B
	// case1. 数据为空，将num存入B中
	// case2. 如果num <= 中位数，将num放入B中
	if minQ.Len() == 0 || num <= minQ.IntSlice[0] {
		heap.Push(minQ, num)
		// 如果A中元素小于B中元素数量-1，说明此时B的中位数已经过大，需要将B的中位数移动到A中
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, heap.Pop(minQ))
		}
		return
	}
	// 如果num > 中位数，将num放入A中
	heap.Push(maxQ, num)
	// maxQ = [2,3], minQ = [1]
	// 当A的长度大于B的长度时，将A中最小的元素放入B中
	if minQ.Len() < maxQ.Len() {
		heap.Push(minQ, heap.Pop(maxQ))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.B.Len() == 0 {
		return 0
	}
	// 偶数个元素：中位数为A的头元素与B的头元素的平均值
	if this.A.Len() == this.B.Len() {
		return float64(this.A.IntSlice[0]+this.B.IntSlice[0]) / 2
	}
	// 奇数个元素：中位数为B的头元素
	return float64(this.B.IntSlice[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such :
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
type minHp struct{ sort.IntSlice }

func (h *minHp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *minHp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type maxHp struct{ sort.IntSlice }

func (h *maxHp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *maxHp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *maxHp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
