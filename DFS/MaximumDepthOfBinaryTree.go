package DFS

import "math"

func maxDepth(root *MinTreeNode) int {
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

	return int(math.Max(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1

}
