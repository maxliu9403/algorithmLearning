package addTwoNumbers

func addTwoNumbersMax(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	// 记录进位
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 求和，如果进位sum=0
		sum := n1 + n2 + carry
		// 计算进位
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 判断最后是否还需进位
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
