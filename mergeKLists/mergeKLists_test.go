package mergeKLists

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//[[1,4,5],[1,3,4],[2,6]]
	nodes := transfer2LisNode([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	v := mergeKLists(nodes)
	fmt.Println("resp:", parse2Nums(v))
}
