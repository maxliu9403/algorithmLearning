package deleteDuplicatesII

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 哑节点，指向head，因为头节点不一定保留，最终返回dummy.Next
	dummy := &ListNode{Next: head}
	cur := dummy
	// 至少要有两个节点才会出现重复节点
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			// 记录此时的重复元素
			x := cur.Next.Val
			// 出现重复元素 开启循环删结点
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			// 没有重复元素，移动当前节点
			cur = cur.Next
		}
	}
	return dummy.Next
}

/*
class Solution:
    def deleteDuplicates(self, head):
        dummy = ListNode(0)
        dummy.next = head
        val_list = []
        while head:
            val_list.append(head.val)
            head = head.next
        counter = collections.Counter(val_list)
        head = dummy
        while head and head.next:
            if counter[head.next.val] != 1:
                head.next = head.next.next
            else:
                head = head.next
        return dummy.next
*/

// 计数
func deleteDuplicatesMethod(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 哑节点，指向head，因为头节点不一定保留，最终返回dummy.Next
	dummy := &ListNode{Next: head}
	filterMap := map[int]int{}
	for head != nil {
		filterMap[head.Val]++
		head = head.Next
	}
	cur := dummy
	for cur.Next != nil {
		count, _ := filterMap[cur.Next.Val]
		if count != 1 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
