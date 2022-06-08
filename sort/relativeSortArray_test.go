package sort

import (
	"log"
	"testing"
)

func TestRelativeSortArrayMethod1(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	arr1 = RelativeSortArrayMethod1(arr1, arr2)
	log.Fatal(arr1)
}

func TestRelativeSortArrayMethod2(t *testing.T) {
	arr1 := []int{19, 2, 3, 1, 3, 2, 4, 6, 7, 9, 2}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	arr1 = RelativeSortArrayMethod2(arr1, arr2)
	log.Fatal(arr1)
}