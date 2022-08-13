package jianzhiOffer

// 字符串中的变位词
func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		// 统计子字符各字符出现次数
		cnt[ch-'a']--
	}
	// 定义左指针
	left := 0
	// 移动右指针
	for right := 0; right < m; right++ {
		// 取出右指针指向的词
		x := s2[right] - 'a' // 计算索引位置，保证是同一个位置且进行抵消
		// 抵消s1中字符的个数，如果是变位词，那么抵消后应该为0
		cnt[x]++
		// 大于0，说明此时s2的left -> right 之间的不是s1的变位词，因为字符数量不一致
		for cnt[x] > 0 {
			// 此时移动左指针，同时将左边元素移除
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
