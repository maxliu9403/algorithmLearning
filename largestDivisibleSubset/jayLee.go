package largestDivisibleSubset

import (
	"fmt"
	"sort"
)

/**
 * 给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
 * answer[i] % answer[j] == 0 ，或
 * answer[j] % answer[i] == 0
 * 如果存在多个有效解子集，返回其中任何一个均可。
 */
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums)) // [[1],[1,2],[1,2,4],[1,2,4,8],[1,2,4,16]]
	for i, num := range nums {
		dp[i] = []int{num}
	}

	var maxSize, maxIndex = 1, 0
	for i := 1; i < len(nums); i++ {
		numI := nums[i]
		for j, numJ := range nums[:i] {
			// 1. numI%numJ = 4%3 != 0, 跳过
			// 2. 4%2 = 0, 判断dp[i]的长度是否小于dp[j]的长度+1（这个1是即将加上的numI）；如果是：则将dp[i]更新为：dp[j] + numI
			if numI%numJ == 0 && len(dp[i]) < len(dp[j])+1 {
				var n = []int{numI}
				for _, v := range dp[j] {
					n = append(n, v)
				}
				//n := append(dp[j], numI) // FIXME：为什么不能使用append？？？
				//if i == 4 {
				//	fmt.Println("append:", n)
				//	fmt.Println("=====dp[j+1]", dp[j+1])
				//}
				dp[i] = n
			}
		}
		if maxSize < len(dp[i]) {
			maxSize, maxIndex = len(dp[i]), i
		}
	}
	fmt.Println("dp:", dp)
	return dp[maxIndex]
}
