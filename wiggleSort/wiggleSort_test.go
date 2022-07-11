package wiggleSort

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	nums := []int{1, 3, 2, 2, 3, 1}
	//nums1 := []int{1, 5, 1, 6, 1, 4}
	// [1,5,1,1,6,4] -> [1,1,1,4,5,6] ->
	// 1, 5, 1, 1, 6, 4
	wiggleSort(nums)
	// Output: [2,3,1,3,1,2]
	fmt.Println("output:", nums)
}
