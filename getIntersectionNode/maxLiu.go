package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	filterMap := map[*ListNode]bool{}
	for tem := headA; tem != nil; tem = tem.Next {
		filterMap[tem] = true
	}
	for tem := headB; tem != nil; tem = tem.Next {
		if filterMap[tem] {
			return tem
		}
	}
	return nil
}

/*
解法2：
https://leetcode.cn/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
*/
func getIntersectionNodeMethod(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		// 达到链表尾节点，指向B链表的头
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
