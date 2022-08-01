package jianzhiOffer

/*
数字最高位位于链表开始位置，所以链表的相加得从尾节点开始，同时还得考虑进位问题

*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 新链表
	l1 = reverse(l1)
	l2 = reverse(l2)
	dummy := &ListNode{}
	sumNode := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 求和，如果进位sum=0
		sum := n1 + n2 + carry
		// 计算进位
		sum, carry = sum%10, sum/10
		sumNode.Next = &ListNode{Val: sum}
		sumNode = sumNode.Next
	}
	if carry > 0 {
		sumNode.Next = &ListNode{Val: carry}
	}
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 记录当前节点的下一个节点
		next := curr.Next
		// 当前节点指向前一个节点
		curr.Next = prev
		// 前一个节点移动到当前节点
		prev = curr
		// 移动后节点
		curr = next
	}
	return prev
}
