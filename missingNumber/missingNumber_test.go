package missingNumber

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	v := XorMissingNumber([]int{3, 0, 1})
	fmt.Println(v)
}
