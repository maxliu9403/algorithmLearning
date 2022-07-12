package findLengthOfLCIS

func findLengthOfLCIS(nums []int) int {
	// dp[i]表示以nums[i]结尾的最长连续递增子序列的长度
	dp := make([]int, len(nums))
	// 初始化dp数组，最少为1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	// 根据题意：连续递增子序列的长度，那么就有nums[i-1] <
	// 计算dp[i]，那么只需要先计算出dp[i-1]，而dp[i] = dp[i-1] + 1，只要有nums[i-1]<nums[i]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
	}
	ans := 0
	// 找出最大长度
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
