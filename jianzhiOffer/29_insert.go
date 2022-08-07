package jianzhiOffer

/*
 1. 大于前节点小于后节点，插入到其中
 2. 如果原始节点长度为0，则结果就是一个节点，那么next指向自己的
 3. 如果原始节点长度1，则将相互指就行
 4. 如果最终都没找到合适的位置，那么就是大于最大值，或者小于最小值，插入到最大值和最小值直接即可
*/

func insert(head *ListNode, x int) *ListNode {
	node := &ListNode{Val: x}
	// 空的
	if head == nil {
		head = node
		head.Next = head
		// 只有一个节点
	} else if head.Next == head {
		head.Next = node
		node.Next = head
		// 两个以上的节点
	} else {
		insertCore(head, node)
	}
	return head
}

func insertCore(head *ListNode, node *ListNode) {
	cur := head
	next := head.Next
	biggest := head
	// 寻找相邻节点插入，next != head 表示循环一圈
	for !(cur.Val <= node.Val && node.Val <= next.Val) && next != head {
		cur = next
		next = next.Next
		// 记录最大值
		if cur.Val >= biggest.Val {
			biggest = cur
		}
	}
	// 插入
	if cur.Val <= node.Val && node.Val <= next.Val {
		cur.Next = node
		node.Next = next
	} else {
		node.Next = biggest.Next
		biggest.Next = node
	}
}
