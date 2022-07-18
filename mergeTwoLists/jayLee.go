package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode
	lessF := func() int {
		if list1 != nil && list2 != nil {
			n1, n2 := list1.Val, list2.Val
			if n1 > n2 {
				list2 = list2.Next
				return n2
			} else {
				list1 = list1.Next
				return n1
			}
		} else if list1 != nil {
			n1 := list1.Val
			list1 = list1.Next
			return n1
		} else {
			n2 := list2.Val
			list2 = list2.Next
			return n2
		}
	}
	for list1 != nil || list2 != nil {
		less := lessF()
		if head == nil {
			head = &ListNode{Val: less}
			tail = head
		} else {
			tail.Next = &ListNode{Val: less}
			tail = tail.Next
		}
	}
	return head
}
