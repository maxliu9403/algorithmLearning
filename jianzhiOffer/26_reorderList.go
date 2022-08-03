package jianzhiOffer

/*
1. 找到链表中点，切割为两个链表
2. 将原链表的左半部分反转
3. 将原链表的两端合并
*/

// 找到中点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转
func reverseListMe(head *ListNode) *ListNode {
	// prev在前，curr在后
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 交叉合并两个链表
func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		// 记录当前当前节点位置
		l1Tmp = l1.Next
		l2Tmp = l2.Next
		// l1的下一个节点指向l2
		l1.Next = l2
		// 移动l1的当前节点
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mid := middleNode(head)
	// 原始链表左边部分
	l1 := head
	// 原始链表右边部分
	l2 := mid.Next
	// 中间断开
	mid.Next = nil
	l2 = reverseListMe(l2)
	mergeList(l1, l2)
	return head
}
