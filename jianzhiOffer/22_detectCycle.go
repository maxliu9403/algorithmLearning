package jianzhiOffer

/*
给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。
*/

/*
  解法：双指针
  设置fast、slow两个指针，fast一次走2步，slow一次走一步
  令：链表总共有a+b个节点，其中 链表头部到链表入口有a个节点（不计链表入口节点)，链表环有b个节点
  第一次相遇：
     令：1. slow走的步数为s，则f = 2s, 一式
         2.fast比slow多走n个环的长度，即f = nb + s 二式
		3. 二式 - 一式 = 0 = nb - s ,即 s = nb，
	也就是说第一次相遇时，slow走的步数是nb ,fast走的步数是2nb
  分析：如果让指针从头节点出发，记录步数k，那么所有走到链表入口节点时的步数:k = a + nb，先走 aa 步到入口节点，之后每绕 1 圈环（ b 步）都会再次到入口节点。经过第一次相遇slow节点已经走了nb步，那么只需要在走a步就到了入口节点。
   此时，可以再次借助双指针，一个节点从头节点出发走a步，slow节点从当前位置也走a步，那么再次相遇的节点必然是入口节点
  第二次相遇：
     1. slow节点位置不变，fast节点指向头节点，slow 和 fast 同时每轮向前走1步；
     此时fast = 0，slow = nb
     2. 当 fast 指针走到 f = a 步时，slow 指针走到步 s=a+nb ，此时 两指针重合，并同时指向链表环入口 。

  复杂度：
	时间复杂度 O(N) ：第二次相遇中，慢指针须走步数 a < a+b；第一次相遇中，慢指针须走步数 a + b - x <a+b，其中 x 为双指针重合点与环入口距离；因此总体为线性复杂度。
空间复杂度 O(1) ：双指针使用常数大小的额外空间。

*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		// slow走一步
		slow = slow.Next
		// 没有环
		if fast.Next == nil {
			return nil
		}
		// fast走两步
		fast = fast.Next.Next
		// 第一次相遇
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			// 第二次相遇
			return p
		}
	}
	return nil
}

/*
  解法2：map
  时间复杂：o(n),空间复杂度o(n)
*/
func detectCycleWithMap(head *ListNode) *ListNode {
	filterMap := map[*ListNode]struct{}{}
	for head != nil {
		if _,ok := filterMap[head]; ok {
			return head
		}
		filterMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}


/*
  解法3：计算出环节点数，
*/