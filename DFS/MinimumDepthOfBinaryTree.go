package DFS

import (
	"math"
)

/*
	二叉树的最小深度
	深度优先解法
	时间复杂度：O(n)
	空间复杂度：O(n)迭代的栈上空间
*/

// 定义二叉数
type MinTreeNode struct {
	Left  *MinTreeNode
	Right *MinTreeNode
	Value int
}

func minDepth(root *MinTreeNode) int {
	/*
			三种情况：
		    1> 如果左右节点都为空，说明到达叶子节点返回0

			2> 如果左节点和右节点其中一个为空，那么需要返回大的那个节点的深度

			3> 如果左右节点都不为空，返回最小深度+1
	*/

	if root == nil {
		return 0
	}

	// 1> 如果左右节点都为空，说明到达叶子节点返回0
	if root.Right == nil && root.Left == nil {
		return 1
	}

	// 3> 如果左右节点都不为空，返回最小深度+1
	return int(math.Min(float64(minDepth(root.Right)), float64(minDepth(root.Left)))) + 1

}
