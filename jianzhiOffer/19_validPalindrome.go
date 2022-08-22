package jianzhiOffer

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// 双指针判断是否是回文
		if s[left] == s[right] {
			left++
			right--
		} else {
			// 删除右指针，判断是否满足回文
			flag1, flag2 := true, true
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			// 删除左指针，判断是否满足回文
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			// 因为是最多删除一个元素，则
			return flag1 || flag2
		}
	}
	return true
}
