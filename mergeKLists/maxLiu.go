package mergeKLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}

/*
	方法1：遍历转换为合并两个有序链表
*/
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	for i := 0; i < len(lists); i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = mergeTwoLists(pre, cur)
	}
	return pre
}

/*
	方法2：对方法1进行优化，采用分而治之
*/
func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	// 计算中点
	mid := left + (right-left)/2
	// 合并左边链表
	leftList := merge(lists, left, mid)
	// 合并右边链表
	rightList := merge(lists, mid+1, right)
	return mergeTwoLists(leftList, rightList)
}

func mergeKListsMethod2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}
