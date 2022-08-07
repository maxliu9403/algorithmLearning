package addTwoNumbers

func addTwoNumbersMax(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	sumNode := dummy
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
		sumNode.Next = &ListNode{Val: sum}
		sumNode = sumNode.Next
	}
	// 判断最后一位是否需要进位
	if carry > 0 {
		sumNode.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
