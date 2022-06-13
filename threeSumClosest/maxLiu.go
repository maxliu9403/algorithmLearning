package threeSumClosest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	var ans = math.MaxInt32
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举
		start, end := i+1, n-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			if absValue(sum-target) < absValue(ans-target) {
				ans = sum
			}
			if sum > target {
				// 去掉重复值，减少遍历次数
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				end--
			} else {
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				start++
			}
		}
	}
	return ans
}

func absValue(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 复杂度分析
// 时间复杂度：O(n^2) 数组排序O(NlogN)，遍历数组O(n)，双指针遍历O(n) 总体:O(NlogN) + O(n)*O(n)
// 空间复杂度：O(1)
