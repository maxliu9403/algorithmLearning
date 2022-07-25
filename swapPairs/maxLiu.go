package swapPairs

type ListNode struct {
	val  int
	Next *ListNode
}

/*
	方法1：遍历
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	// 初始化交换2个节点的前一个节点
	temp := dummy
	// 至少要有两个节点
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		// 2号节点交换到1号位置
		temp.Next = node2
		// 1号节点的下一个节点指向2号节点的下一个节点
		node1.Next = node2.Next
		// 2号节点的下一个节点指向1号节点
		node2.Next = node1
		// 更新初始节点指向
		temp = node1
	}
	return dummy.Next
}
