package insertionSortList

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	list := trans2ListNode([]int{4, 2, 1, 3})
	v := insertionSortList(list)
	fmt.Println(trans2Slice(v))
}
