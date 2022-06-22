package insertionSortList

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先转化为切片，排序后再转化为ListNode
func insertionSortList(head *ListNode) *ListNode {
	arr := trans2Slice(head)
	sort.Ints(arr)
	return trans2ListNode(arr)
}

func trans2Slice(head *ListNode) []int {
	var resp []int
	for head != nil {
		resp = append(resp, head.Val)
		head = head.Next
	}
	return resp
}

func trans2ListNode(arr []int) *ListNode {
	h := &ListNode{}
	// 重新定义一个链表，使得h的位置变更不会影响到res。
	// 对h指向的目标的修改影响到了res指向的目标。
	res := h
	for _, val := range arr {
		h.Next = &ListNode{Val: val}
		h = h.Next
	}
	return res.Next
}
