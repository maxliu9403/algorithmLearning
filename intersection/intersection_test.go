package intersection

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	v := intersection([]int{1, 1, 2, 2}, []int{2, 2})
	fmt.Println("resp:", v)
}
