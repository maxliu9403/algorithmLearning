package BFS

func maxDepth(root *MaxTreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*MaxTreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) != 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
