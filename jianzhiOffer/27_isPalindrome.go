package jianzhiOffer

/*
 如果是回文链表，那么从中点切开，翻转一侧后应该是一致的
思路：
	1. 找到链表中点，切开
    2. 翻转一侧
    3. 判断是否一致
*/

func middleNodeMe(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseListMethod(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	mid := middleNodeMe(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseListMethod(l2)
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
