package deleteDuplicatesII

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	tail := dummyHead
	// [null,1,2,3,3,4,4,5]
	for tail.Next != nil && tail.Next.Next != nil {
		if tail.Next.Val == tail.Next.Next.Val {
			v := tail.Next.Val
			tail.Next = tail.Next.Next.Next
			// 2个以上相同时，一样会有问题：[1,1,1,2,3]
			for tail.Next != nil && tail.Next.Val == v {
				tail.Next = tail.Next.Next
			}
		} else {
			tail = tail.Next
		}
	}
	return dummyHead.Next
}
