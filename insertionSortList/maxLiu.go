package insertionSortList

type ListNodeMax struct {
	Val  int
	Next *ListNodeMax
}

func insertionSortListMax(head *ListNodeMax) *ListNodeMax {
	dummyHead := &ListNodeMax{Val: 0}
	dummyHead.Next = head
	cur := head
	var prev *ListNodeMax
	var temp *ListNodeMax

	for cur != nil && cur.Next != nil { // cur指针扫整个链表
		if cur.Val <= cur.Next.Val { // 符合递增，继续推进
			cur = cur.Next
		} else { // 找到需要变动的cur.Next
			temp = cur.Next          // 保存给temp
			cur.Next = cur.Next.Next // 删除结点

			prev = dummyHead                // 从dummy开始扫，用prev推进，找插入的位置
			for prev.Next.Val <= temp.Val { // 继续推进
				prev = prev.Next
			}
			// 此时prev.Next.Val更大，插入到 prev 和 prev.Next 之间
			temp.Next = prev.Next
			prev.Next = temp // 先改temp.Next，再接给prev.Next
		}
	}

	return dummyHead.Next // 就算头结点发生改变了，也能通过dummyHead.Next获取到头结点
}
