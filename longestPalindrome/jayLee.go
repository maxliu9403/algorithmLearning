package longestPalindrome

/**
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 * 回文串: 一个正读和反读都一样的字符串
 * 状态转移方程：dp[i,j] = dp[i+1,j-1] && s[i] = s[j]
 */
func longestPalindrome(s string) string {
	var res string
	dp := make([][]bool, len(s)) // dp设置为一个横竖各位len(s)长度的二维数组，dp[i][j]代表 Si ~ Sj 是否为回文子串
	// 长度为1时, 自己和自己比较 一定是 true, 其余置为 false
	for i := range s {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		res = s[i : i+1] // 回文串: s[i]
	}
	// 长度为2时, 自己和下一位比较 相同为 true, 其余为 false
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == string(s[i+1]) {
			dp[i][i+1] = true
			res = s[i : i+2] // 回文串: s[i] + s[i+1]
		}
	}
	// 长度大于2时, 判断DP[i+1, j-1]是否为回文串，并且s[i]和s[j]是否相同，都满足为 true, 其余为 false
	for length := 3; length <= len(s); length++ {
		for i := 0; i+length <= len(s); i++ {
			j := i + length - 1
			// 状态转移方程
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				if len(s[i:j+1]) > len(res) {
					res = s[i : j+1] // 更新回文串
				}
			}
		}
	}
	return res
}
