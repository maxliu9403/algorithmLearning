package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

// CommonReverseList 反转链表
/**
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		// curr的下一位指向prev
		curr.Next = prev
		// prev和curr各右移一位
		prev, curr = curr, next
	}
	return prev
}
