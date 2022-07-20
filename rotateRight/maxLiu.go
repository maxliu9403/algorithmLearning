package rotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	思路：构建环
	1 -> 2 -> 3 -> 4 -> 5, k=4
	5 -> 1 -> 2 -> 3 -> 4
	4 -> 5 -> 1 -> 2 -> 3
	3 -> 4 -> 5 -> 1 -> 2
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	// 计算链表的长度
	length := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		length++
	}
	// 计算需要移动多少次
	add := length - k%length
	// 移动次数是节点总数的整数备
	if add == length {
		return head
	}
	// 最后一个节点指向头节点
	iter.Next = head
	// 移动节点
	for add > 0 {
		iter = iter.Next
		add--
	}
	// 移动节点后的下一个节点就是最终结果
	ret := iter.Next
	// 断开
	iter.Next = nil
	return ret
}
