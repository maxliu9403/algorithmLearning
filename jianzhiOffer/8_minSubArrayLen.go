package jianzhiOffer

import "math"

/*
 滑动窗口
时间复杂度：O(n)，其中 nn 是数组的长度。指针 left 和 right 最多各移动 n 次。
空间复杂度：O(1)
*/
func minSubArrayLen(target int, nums []int) int {
	// 定义左指针
	left, sum := 0, 0
	minLength := math.MaxInt32
	n := len(nums)
	// 从0下标开始遍历右指针
	for right := 0; right < n; right++ {
		// 累计
		sum += nums[right]
		// 如果累计值大于目标值，通过移动左指针减少长度
		for left <= right && sum >= target {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
