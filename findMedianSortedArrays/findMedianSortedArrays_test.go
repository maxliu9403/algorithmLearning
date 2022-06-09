package findMedianSortedArrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	// 0 = 0
	// 0,1,2 = 1
	// 0,1,2,3,4 = 2
	// 0,1,2,3,4,5,6,7 = 3
	fmt.Println("======奇数中位数索引", 7/2)

	// 0,1 = 1
	// 0,1,2,3 = 2
	// 0,1,2,3,4,5 = 3
	fmt.Println("======偶数中位数索引", (4/2)-1, 4/2)

	v := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println("====value", v)
}
