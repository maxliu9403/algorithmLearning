package labuladong

/*
思路：用左右两个指针遍历s字符串，当滑动窗口中的字符不能覆盖t中的字符时，右指针右移，扩大窗口，把右边的字符加入滑动窗口，当滑动窗口中的字符能覆盖t中的字符时，不断左移左指针，缩小窗口，直到窗口中的字符刚好能覆盖t中的字符，这个时候在左移就不能覆盖t中的字符了，在指针移动的过程中，不断更新最小覆盖子串
复杂度：时间复杂度o(n)，n是s的长度，空间复杂度o(t)，t是字符集的大小
*/
func minWindow(s string, t string) string {
	// need 记录子串中各字符的个数
	// window 记录窗口中各字符出现的次数
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// 初始化左右窗口边界
	left, right := 0, 0
	// 记录最小覆盖子串的启始索引及最小长度，加一是为了考虑s本身就是最小子串
	start, end, subStrShortestLen := 0, 0, len(s)+1
	// 记录总共有多少个字符满足条件了
	valid := 0
	for right < len(s) {
		// 将要移入窗口的字符
		c := s[right]
		// 移动右窗口
		right++
		// 判断该元素是否是最小覆盖子串中的字符
		_, ok := need[c]
		if ok {
			// 窗口记录增加
			window[c] ++
			// 判断窗口中记录的字符个数是否等于需要的个数
			// 字符出现次数等于need中该字符出现的次数，则记录
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左窗口是否需要收缩，因为是寻早最小子串
		// valid 总共有多少个字符满足了条件，len(need)需要多少个满足条件
		for valid == len(need) {
			// 判断此时更新得倒的窗口是否比上次更新的窗口还要小，如果是才更新，这样保证最终退出时的结果一定是最小的
			if right-left < subStrShortestLen {
				start, end = left, right
				subStrShortestLen = right - left
			}
			d := s[left]
			// 右移左边界
			left++
			// 判断此时的元素是否是目标元素
			_, ok = need[d]
			if ok {
				// 判断目标数量和窗口中数量是否一致，如果一致计数器减去
				if window[d] == need[d] {
					// 因为valid是只要满足
					valid--
				}
				// 减少窗口中记录的元素个数，直到恰好等于目标数量
				window[d]--
			}
		}
	}
	return s[start:end]
}
