package singleNumber

import (
	"fmt"
	"testing"
)

func TestXorSingleNumber(t *testing.T) {
	v := XorSingleNumber([]int{3, 5, 1, 1, 2, 2})
	fmt.Println("resp", v)
}
