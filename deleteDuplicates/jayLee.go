package deleteDuplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}
	return head
}
