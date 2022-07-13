package findLengthOfLCIS

func dpFindLengthOfLCIS(nums []int) int {
	dp := make([][]int, len(nums))
	for i, num := range nums {
		dp[i] = []int{num}
	}
	for i := 1; i < len(nums); i++ {
		// 倒序遍历
		j := i - 1
		if nums[j] < nums[i] && len(dp[i]) < len(dp[j])+1 {
			dp[i] = append(dp[j], nums[i])
		}
	}
	var maxSize int
	for _, value := range dp {
		if maxSize < len(value) {
			maxSize = len(value)
		}
	}
	return maxSize
}
