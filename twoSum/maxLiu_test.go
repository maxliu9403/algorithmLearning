package twoSum

import (
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 2, 2, 1, 1}
	ans := twoSumTarget(nums, 3)
	log.Fatal(ans)
}
