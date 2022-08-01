package jianzhiOffer

func reverseList(head *ListNode) *ListNode {
	// prev在前，curr在后
	var prev *ListNode
	curr := head
	for curr != nil {
		// 保存当前节点的下一个节点
		next := curr.Next
		// 让curr的下一个节点指向prev，实现局部翻转
		curr.Next = prev
		// 移动prev到下一个节点
		prev = curr
		// 移动curr
		curr = next
	}
	return prev
}
