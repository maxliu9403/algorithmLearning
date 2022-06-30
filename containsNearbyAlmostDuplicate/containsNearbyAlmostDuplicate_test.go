package containsNearbyAlmostDuplicate

import (
	"fmt"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	v := containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	fmt.Println(v)
}
