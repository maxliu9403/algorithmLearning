package lengthOfLIS

/*
思路：动态规划
步骤：
	1. dp[i] 表示以nums[i]结尾的最长递增子序列
	2. 求解 dp[i] 时，向前遍历找出比 i 元素小的元素 j，令 dp[i] 为 max（dp[i],dp[j]+1)
时间复杂度：o(n^2)
空间复杂度：o(n)
*/

func lengthOfLIS(nums []int) int {
	// dp[i] 表示以nums[i]结尾的最长递增子序列
	dp := make([]int, len(nums))
	// 初始化dp数组
	for i:=0; i <len(nums); i ++ {
		dp[i]= 1
	}
	// 计算dp[i]，只需要
	for i := 0; i<len(nums); i ++ {
		for j :=0; j<i;j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for i:=0; i<len(nums); i++ {
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
