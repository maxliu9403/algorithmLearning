package mergeTwoLists

func mergeTwoListsMax(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		mergeTwoLists(list2.Next, list1)
		return list2
	}
}
