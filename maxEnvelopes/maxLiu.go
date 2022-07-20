package maxEnvelopes

import (
	"sort"
)

/*
	思路：按照高度升序排，对宽转换为子长子序列问题。对于相同高时对宽进行降序排，这里保证在计算最长子序列时候避免结果集中包含相同的高
*/
func MaxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	// 转换为子长子序列问题
	return lengthOfLIS(envelopes)
}

func lengthOfLIS(envelopes [][]int) int {
	// dp[i] 表示以nums[i]结尾的最长递增子序列
	dp := make([]int, len(envelopes))
	// 初始化dp数组
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
	}
	// 计算dp[i]，只需要
	for i := 0; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
