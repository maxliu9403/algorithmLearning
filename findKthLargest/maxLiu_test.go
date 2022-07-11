package findKthLargest

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3,2,1,5,6,4}
	QuickSort(nums, 0, 5)
	log.Fatal(nums)
}
