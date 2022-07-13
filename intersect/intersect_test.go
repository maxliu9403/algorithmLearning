package intersect

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	v := intersect([]int{1, 1, 2, 2}, []int{2, 2, 1})
	fmt.Println("===reps :", v)
}
