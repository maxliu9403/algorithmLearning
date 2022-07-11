package topKFrequent

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	v := topKFrequent([]int{3, 0, 1, 0}, 1)
	fmt.Println("===========", v)
}
