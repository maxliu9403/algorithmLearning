package majorityElement2

import (
	"fmt"
	"testing"
)

// [1, 2, 3, 4] => []，cd1 = 4, ct1 = 1, cd2 = 2, ct2 = 0
// [1, 1, 2, 3] => [1], cd1 = 1, ct1 = 1, cd2 = 2, ct2 = 0
// [1, 1, 2, 2] => [1, 2], cd1 = 1, ct1 = 2, cd2 = 2, ct2 = 2
func TestMajorityElement2(t *testing.T) {
	v := MajorityElement([]int{2, 2, 1, 3})
	fmt.Println(v)
}
