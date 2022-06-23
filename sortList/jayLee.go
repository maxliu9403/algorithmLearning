package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 自底向上 - 递归
func sortListBottomToUpWithRecursion(head *ListNode) *ListNode {
	return toSortList(head, nil)
}

// TODO 自底向上 - 轮询
func sortListBottomToUpWithLoop(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	resp := &ListNode{Next: head}
	l := length(head)
	for subLength := 1; subLength < l; subLength <<= 1 {
		prev, cur := resp, resp.Next
		for cur != nil {
			h1 := cur // 左半区间的头指针
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			h2 := cur.Next // 右半区间的头指针
			cur.Next = nil
			cur = h2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode // 合并左右半区后的下一个指针
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge(h1, h2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return resp.Next
}

func toSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil // 最后一位，需要断开与另一区间的联系
		return head
	}
	// 中间位置
	mid := middle(head, tail)
	return merge(toSortList(head, mid), toSortList(mid, tail)) // 区分区间，递归合并
}

// 找中间值
func middle(head *ListNode, tail *ListNode) *ListNode {
	slow, fast := head, head
	// slow一次移动一位，fast一次移动两位。当fast移动到最后一位时，slow就是中间位置。
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return slow
}

// merge: 合并2个有序的链表
func merge(h1, h2 *ListNode) *ListNode {
	resp := &ListNode{}
	t, t1, t2 := resp, h1, h2
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			t.Next = t1
			t1 = t1.Next
		} else {
			t.Next = t2
			t2 = t2.Next
		}
		t = t.Next
	}
	if t1 != nil {
		t.Next = t1
	} else if t2 != nil {
		t.Next = t2
	}
	return resp.Next
}

// length: 链表长度
func length(h *ListNode) int {
	var length int
	for l := h; l != nil; l = l.Next {
		length++
	}
	return length
}
