package removeDuplicates2

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesMedium(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6}
	a := RemoveDuplicatesMedium(nums)
	fmt.Println(a)
}

func TestRemoveDuplicatesK(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6}
	a := RemoveDuplicatesK(nums, 2)
	fmt.Println(a)
}