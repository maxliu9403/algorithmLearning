package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 */
func swapPairs(head *ListNode) *ListNode {
	// p: 1 -> 3 -> 4
	// pre: 0 -> 1 -> 2 -> 3 -> 4
	h := &ListNode{Next: head}
	pre, p := h, head
	for p != nil && p.Next != nil {
		pre.Next, p.Next = p.Next, p.Next.Next
		pre.Next.Next, p = p, p.Next
		pre = pre.Next.Next
	}
	return h.Next
}
