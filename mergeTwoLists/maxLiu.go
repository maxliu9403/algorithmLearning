package mergeTwoLists

func mergeTwoListsMax(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}

func mergeTwoListsMethod2(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			// 单前节点指针指向值小的节点
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			// 同时移动当前节点
			cur = cur.Next
			// 通过上面的判断，至多存在一个链表为非空，而链表都是有序的，所以非空链表的一定大于当前节点，只需要把当前节点指向非空链表即可
		} else if list1 != nil {
			cur.Next = list1
			break
		} else {
			cur.Next = list2
			break
		}
	}
	return head.Next
}
