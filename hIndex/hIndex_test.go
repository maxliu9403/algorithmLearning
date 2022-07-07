package hIndex

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	v := hIndex([]int{11, 12})
	fmt.Println("H指数：", v)
}
