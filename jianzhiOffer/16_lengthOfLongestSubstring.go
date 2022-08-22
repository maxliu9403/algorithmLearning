package jianzhiOffer

func lengthOfLongestSubstring(s string) int {
	ans := 0
	if len(s) == 0 {
		return ans
	}
	n := len(s)
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	left := -1
	right := 0
	// 记录最长子串中是否有重复字符
	countDump := 0
	for right = 0; right < n; right++ {
		// 记录每个字符出现次数
		m[s[right]]++
		// 出现重复字符，移动左边界
		if m[s[right]] == 2 {
			countDump++
		}
		// 判断此时的子串是否还有重复字符
		for countDump > 0 {
			left++
			m[s[left]]--
			if m[s[left]] == 1 {
				countDump--
			}
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
