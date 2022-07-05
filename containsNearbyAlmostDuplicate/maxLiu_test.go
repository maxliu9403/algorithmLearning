package containsNearbyAlmostDuplicate

import (
	"log"
	"testing"
)

func TestContainsNearbyAlmostDuplicate2(t *testing.T) {
	ContainsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0)
}

func TestQuickSort(t *testing.T) {
	nums := []int{1,5,9,1,5,9}
	QuickSort(nums, 0, 5)
	log.Fatal(nums)
}