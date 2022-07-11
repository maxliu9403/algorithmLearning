package sortList

func sortList(head *ListNode) *ListNode {
	// 递归的出口，不用排序 直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢指针
	fast, slow := head, head
	var preSlow *ListNode
	// 判断快指针是否已经走到头
	for fast != nil && fast.Next != nil {
		preSlow = slow
		// 慢指针走一步，快指针走两步
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 断开链表
	if preSlow != nil {
		preSlow.Next = nil
	}
	l := sortList(head) // 已排序的左链
	r := sortList(slow) // 已排序的右链
	return mergeList(l, r)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	prev := dummy
	// 左右链表元素都存在
	if l1 != nil && l2 != nil {
		// 左链表元素小于右链表元素，
		if l1.Val < l2.Val {
			// 指向小的元素
			prev.Next = l1
			// 判断下一个元素
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
	}
	// l2没有元素
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	// 真实头结点
	return dummy.Next
}
