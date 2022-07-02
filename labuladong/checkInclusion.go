package labuladong

//给你两个字符串s1和s2 ，写一个函数来判断s2是否包含 s1的排列。如果是，返回true；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。

/*
思路：由于排列不会改变字符串中每个字符的个数，所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的排列
*/

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	// 记录总共有多少个字符满足条件了
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		_, ok := need[c]
		if ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左窗口是否需要更新
		// 因为各种排列的长度显然应该是一致的
		for right-left == len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			_, ok = need[d]
			if ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
