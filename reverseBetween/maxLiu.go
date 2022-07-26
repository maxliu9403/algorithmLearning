package reverseBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{Next: head}
	// 找到开始反转开始节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 从pre出发找到反转链表的结束节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 截取子链表
	leftNode := pre.Next
	// 记录子链表的结束节点的下一个节点
	cur := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil
	// 反转子链表
	reverseList(leftNode)

	// 拼接
	pre.Next = rightNode
	leftNode.Next = cur
	return dummy.Next
}

/*
 curr：指向待反转区域的第一个节点 left
 next: 永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
 pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变。
*/
func reverseBetweenMethod(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	// 找到开始反转开始节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
