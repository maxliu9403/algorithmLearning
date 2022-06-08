package twoSum

import (
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	ans := TwoSum(nums, 9)
	log.Fatal(ans)
}
