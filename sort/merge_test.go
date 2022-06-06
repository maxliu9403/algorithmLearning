package sort

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := Merge(intervals)
	log.Fatal(res)
}
