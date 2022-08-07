package jianzhiOffer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getCount(head *ListNode) int {
	var count int
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

func getIntersectionNodeMethod(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil  {
		return nil
	}
	// 计算两个链表的长度
	countA := getCount(headA)
	countB := getCount(headB)

	if countA > countB {
		// 选出长的链表，先移动
		gap := countA - countB
		for i := 0; i < gap; i++ {
			headA = headA.Next
		}
	} else {
		gap := countB - countA
		for i := 0; i < gap; i++ {
			headB = headB.Next
		}
	}
	// 长和短链表同时移动，遇到相同就退出
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
