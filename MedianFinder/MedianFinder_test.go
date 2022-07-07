package MedianFinder

import (
	"fmt"
	"testing"
)

/**
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
*/
func TestMedianFinder(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	obj := Constructor()
	for _, num := range nums {
		obj.AddNum(num)
		param := obj.FindMedian()
		fmt.Println("===", param) // 1.5
	}
}
