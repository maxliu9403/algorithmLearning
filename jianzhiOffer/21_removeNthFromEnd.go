package jianzhiOffer

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	front := head
	back := dummy
	// 将前节点移动到n-1节点位置
	for i := 0; i < n; i++ {
		front = front.Next
	}
	// 同时移动前后节点，此时前后节点相查n个节点
	for front != nil {
		front = front.Next
		back = back.Next
	}
	// 前节点到达尾节点，此时后节点正好在倒数第k+1节点，将倒数k+1节点指向k-1节点，这样就删除倒数k节点
	back.Next = back.Next.Next
	return dummy.Next
}

/*
两次遍历
注意不能直接在removeNthFromEndMethod2计算链表长度，因为会改变头节点的位置
*/
func getLength(head *ListNode, length int) int {
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

func removeNthFromEndMethod2(head *ListNode, n int) *ListNode {
	// 第一次遍历计算链表长度
	length := 0
	length = getLength(head, length)
	dummy := &ListNode{Next: head}
	cur := dummy
	// 找到第倒数k-1节点
	for i:=0; i<length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
