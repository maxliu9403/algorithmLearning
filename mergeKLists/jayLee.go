package mergeKLists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if 0 == len(lists) {
		return nil
	}

	iHeap := new(IHeap)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(iHeap, lists[i]) // 依次插入队列，按照当前值最小值排序
		}
	}

	head := new(ListNode)
	tail := head
	for iHeap.Len() != 0 {
		tail.Next = heap.Pop(iHeap).(*ListNode) // 推出当前最小值的点
		tail = tail.Next
		if tail.Next != nil {
			heap.Push(iHeap, tail.Next)
		}
	}
	return head.Next
}

type IHeap []*ListNode

func (hp IHeap) Len() int           { return len(hp) }
func (hp IHeap) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }
func (hp IHeap) Less(i, j int) bool { return hp[i].Val < hp[j].Val }

func (hp *IHeap) Push(x interface{}) { *hp = append(*hp, x.(*ListNode)) }
func (hp *IHeap) Pop() interface{}   { n := len(*hp); x := (*hp)[n-1]; *hp = (*hp)[:n-1]; return x }

func transfer2LisNode(t [][]int) []*ListNode {
	var resp []*ListNode
	for _, nums := range t {
		head := new(ListNode)
		tail := head
		for _, num := range nums {
			tail.Next = &ListNode{Val: num}
			tail = tail.Next
		}
		resp = append(resp, head.Next)
	}
	return resp
}

func parse2Nums(t *ListNode) []int {
	var resp []int
	for t != nil {
		resp = append(resp, t.Val)
		t = t.Next
	}
	return resp
}
