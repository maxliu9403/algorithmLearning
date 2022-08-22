package jianzhiOffer

func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	count := 0
	for i := 0; i < n; i++ {
		// 奇数回文串
		count += countPalindrome(s, i, i)
		// 偶数回文串
		count += countPalindrome(s, i, i+1)
	}
	return count
}

func countPalindrome(s string, left, right int) int {
	count := 0
	for left < right && right < len(s) && s[left] == s[right] {
		// 向左右延伸一个元素
		left--
		right++
		count++
	}
	return count

}
