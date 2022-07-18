package isMatch

/**
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 * - '.' 匹配任意单个字符
 * - '*' 匹配零个或多个前面的那一个元素
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 *
 * 分析方法
 * 1 判断是否可以使用动态规划算法: 能够划分为子问题，并且满足"自底向上"
 * 2 确定状态: 定义状态数组: dp[i][j] 代表: s的前i个元素与p的前j个元素的是否匹配
 * 3 子问题:
 * 	3.1 化繁为简: 匹配的逻辑暂时忽略, 用match()代替, 先把整体逻辑撸顺!!
 *  3.2
 */
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
