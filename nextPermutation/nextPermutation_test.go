package nextPermutation

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	v := []int{2, 3, 1, 3, 3, 3}
	nextPermutation(v)
	fmt.Println(v)
}
