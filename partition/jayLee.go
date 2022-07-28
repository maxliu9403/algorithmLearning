package partition

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 */
func partition(head *ListNode, x int) *ListNode {
	leftHead, rightHead := new(ListNode), new(ListNode)
	left, right, tail := leftHead, rightHead, head
	for tail != nil {
		if tail.Val < x {
			left.Next = &ListNode{Val: tail.Val}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: tail.Val}
			right = right.Next
		}
		tail = tail.Next
	}
	left.Next = rightHead.Next
	return leftHead.Next
}
