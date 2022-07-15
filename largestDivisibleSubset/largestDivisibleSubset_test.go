package largestDivisibleSubset

import (
	"fmt"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	v := largestDivisibleSubset([]int{1, 2, 4, 8, 16})
	fmt.Println("v:", v)
}
