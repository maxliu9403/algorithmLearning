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

func reverseListWithRecursion(head *ListNode) *ListNode {
	// base-case
	if head.Next == nil {
		return head
	}
	// head: 1 -> reverse(2 -> 3 -> 4 -> 5 -> null)
	// 结果: last: null <- 2 <- 3 <- 4 <- 5
	// 		head: 1 -> 2 -> null
	last := reverseListWithRecursion(head.Next) // reverse的定义：输入一个节点head，将「 以head为起点 」的链表反转，并返回反转之后的头结点。
	// head: 1 <- 2 <- 3 <- 4 <- 5
	// head: 1 -> 2
	head.Next.Next = head
	head.Next = nil
	return last
}
