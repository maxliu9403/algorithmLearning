package findKthLargest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSearch(t *testing.T) {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}
	rand.Seed(time.Now().UnixNano())
	v := quickSelectRandomizedKth(A, 0, len(A)-1, len(A)-2)
	fmt.Println(A, v)
}

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// RANDOM(5, 15)
	v := rand.Int()%(10) + 5
	fmt.Println(v)
}
