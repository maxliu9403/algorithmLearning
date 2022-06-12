package threeSumClosest

import (
	"log"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ans := ThreeSumClosest(nums, target)
	log.Fatal(ans)
}
