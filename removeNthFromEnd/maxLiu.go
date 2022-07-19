package removeNthFromEnd

/*
方法1：计算链表长度，找到要删除节点的前一个节点，将它的Next指向删除节点后节点
删除节点 = length - n+1
*/

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEndMax(head *ListNode, n int) *ListNode {
	// 不能直接在函数内计算
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	// 修改删除节点的前一个节点的Next
	cur.Next = cur.Next.Next
	return dummy.Next
}

/*
方法2：栈，遍历链表，元素添加到栈（先进后出）通过下标计算出删除节点前一个节点prev
prev.Next = prev.Next.Next
*/
func removeNthFromEndStack(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	// 删除节点前节点
	prev := nodes[len(nodes)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}

/*
方法3：快慢指针
然后先让 fast 往前走 n 步，slow 指针不动，这时候两个指针的距离为 n。
再让 slow 和 fast 同时往前走（保持两者距离不变），直到 fast 指针到达结尾的位置。
这时候 slow 会停在待删除节点的前一个位置，
*/
func removeNthFromEndMethod(head *ListNode, n int) *ListNode {
	// 定义哑节点
	dummy := &ListNode{0, head}
	// 快指针指向头节点
	// 慢指针指向哑节点
	fast, slow := head, dummy
	// 快指针先走n个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
