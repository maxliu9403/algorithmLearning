package fourSum

import (
	"sort"
)

/**
题目：
	给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]]
	（若两个四元组元素一一对应，则认为两个四元组重复）：
	- 0 <= a, b, c, d < n
	- a、b、c 和 d 互不相同
	- nums[a] + nums[b] + nums[c] + nums[d] == target
	- 你可以按 任意顺序 返回答案 。

思路：
	在三数之和基础上多加一层循环即可。
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var response [][]int
	for i := 0; i < len(nums); i++ {
		if 1 <= i && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if i+2 <= j && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if j+2 <= left && nums[left] == nums[left-1] {
					left++
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					response = append(response, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				}
			}
		}
	}
	return response
}
