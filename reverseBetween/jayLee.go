package reverseBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表。
 */
// 迭代法
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// dummyHead = [null,1,2,3,4,5], left = 2, right = 4
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyHead := &ListNode{Next: head}
	// 第 1 步：截取链表
	pre := dummyHead // pre = left 前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	leftNode := pre.Next // leftNode = left节点
	rightNode := pre     // rightNode = right 节点
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	nxt := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	// 第 2 步：反转链表的子区间
	reverse(leftNode) // 2 <- 3 <- 4
	// 第 3 步：接回到原来的链表中 1 -> (4 -> 3 -> 2) -> 5
	pre.Next = rightNode
	leftNode.Next = nxt
	return dummyHead.Next
}

func reverse(head *ListNode) {
	// 记录上一位
	var pre *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
}

// 递归法
func reverseBetweenWithRecursion(head *ListNode, left, right int) *ListNode {
	if left == 1 {
		return reverseNWithRecursion(head, right) // 反转前N个元素
	}
	// head依次右移，目标left和right也就相对应的依次减去1，知道left为1时。问题转化为：反转前N个元素。
	head.Next = reverseBetweenWithRecursion(head.Next, left-1, right-1) // 定义：反转head.Next中[left-1...right-1]元素
	return head
}

// 递归法反转前N个元素
var suc *ListNode // 后置节点
func reverseNWithRecursion(head *ListNode, n int) *ListNode {
	// 1 -> reverse(2 -> 3 -> 4) -> 5
	if n == 1 {
		suc = head.Next // 记录第 n + 1 个节点
		return head
	}
	lst := reverseNWithRecursion(head.Next, n-1) // 定义：反转head.Next中前n-1个元素
	// 1 -> 2 <- 3 <- 4 -> 5
	head.Next.Next = head // 1 <- 2 <- 3 <- 4
	head.Next = suc       // 5 <- 1 <- 2 <- 3 <- 4
	return lst
}
