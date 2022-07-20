package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	// 初始化pre, end都指向dummy
	// pre为每次要翻转链表的头节点的上一个节点
	// end为每次要翻转链表的尾节点
	pre, end := dummy, dummy
	for end.Next != nil {
		// 移动k次end指针，同时保证不要越界
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 判断此时end是否为空，即需要翻转的链表的节点数小于k，此时不需要翻转
		if end == nil {
			break
		}
		// 记录翻转的链表最后节点的下一个节点，方便后面链接链表
		next := end.Next
		// 断开翻转的链表
		end.Next = nil
		// 翻转链表的开始节点
		start := pre.Next
		// 更新pre指向到已翻转的链表
		pre.Next = reverse(start)
		// 翻转链表的开始节点变成结束节点
		start.Next = next
		// 将pre换成下次要翻转的链表的头结点的上一个节点。即start
		pre = start
		// 翻转结束，将end置为下次要翻转的链表的头结点的上一个节点。即start
		end = start
	}
	return dummy.Next

}

func reverse(head *ListNode) *ListNode {
	// 单链表为空或只有一个节点，直接返回原单链表
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode = nil, nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
