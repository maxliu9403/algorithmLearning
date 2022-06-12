package maxArea

import (
	"log"
	"testing"
)

func TestMaxArea(t *testing.T) {
	nums := []int{1,8,6,2,5,4,8,3,7,9}
	ans := MaxArea(nums)
	log.Fatal(ans)
}
