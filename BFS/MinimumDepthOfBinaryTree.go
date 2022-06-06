package BFS

/*
	二叉树的最大深度
	广度优先
	时间复杂度：O(n)
	空间复杂度：O(n)，每层节点的最大数量，最优情况为O(1)
*/

type MaxTreeNode struct {
	Left  *MaxTreeNode
	Right *MaxTreeNode
	Value int
}

func minDepth(root *MaxTreeNode) int {
	if root == nil {
		return 0
	}
	q := []*MaxTreeNode{root}
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth

}
