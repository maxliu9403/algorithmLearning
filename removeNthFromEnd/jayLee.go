package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slowP, fastP := dummyHead, head
	// 总长度l = n + (l - n)，所以可以通过"剔除"前n个结点，获得剩余l-n个的剩余结点
	for n > 0 {
		fastP = fastP.Next
		n--
	}
	// 扫描l - n个剩余指针后，slowP的Next正好为倒数第n个节点
	for fastP != nil {
		slowP, fastP = slowP.Next, fastP.Next
	}
	slowP.Next = slowP.Next.Next
	return dummyHead.Next
}
