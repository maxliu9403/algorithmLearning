package containsDuplicate

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{1,2,3,1}
	QuickSort(nums, 0, 3)
	log.Fatal(nums)
}