package DFS

import (
	"fmt"
	"testing"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	//root := &MinTreeNode{
	//	Value: 2,
	//	Right: &MinTreeNode{Value: 3, Right: &MinTreeNode{Value: 4, Right: &MinTreeNode{Value: 5, Right: &MinTreeNode{Value: 6}}}},
	//}
	root := &MinTreeNode{
		Value: 2,
		Right: &MinTreeNode{Value: 3, Right: &MinTreeNode{Value: 4, Right: &MinTreeNode{Value: 5, Right: &MinTreeNode{Value: 6}}}},
		Left:  &MinTreeNode{Value: 9, Right: &MinTreeNode{Value: 4}, Left: &MinTreeNode{Value: 3}},
	}
	ans := maxDepth(root)
	fmt.Println(ans)

}
