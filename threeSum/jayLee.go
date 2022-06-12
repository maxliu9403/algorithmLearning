package threeSum

import (
	"sort"
)

/**
题目：
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
	请你找出所有和为 0 且不重复的三元组

思路：
	题目转化为：给你一个长度为 n 的整数数组 nums 和 一个目标值 target = 0。请你从 nums 中选出所有不重复的三个整数，使它们的和等于 target，如果没有，返回空数组。
	遍历从小到大排序后的数组 i，放置双指针的位置为：下一位 j 和最后一位 z，求出3个位置的和 sum。
	1. 如果 sum 正好等于 target：由于可能存在重复的三元组，判断该3个元素是否存在，不存在：记录为新三元组；存在：不记录。考虑到可能存在 i 不变， 不止一个 j 和 z 满足要求的场景， 左右指针同时移动一位。
	2. 如果 sum 小于 target：说明整体位置偏左，左指针右移。
	3. 如果 sum 大于 target：说明整体位置偏右，右指针左移。

优化：threeSumUpgrade
	1. 判断三元组是否重复：重复说明有相邻2个元素值相同，因此通过比较是否与前一位元素相同进行过滤。
	2. 优化所有中间数组和函数。
*/
func threeSumUpgrade(nums []int) [][]int {
	// FAST PATH
	if len(nums) < 3 {
		return [][]int{}
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		}
		return [][]int{}
	}

	sort.Ints(nums)
	var response [][]int
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		// 优化1
		if 1 <= i && nums[i] == nums[i-1] {
			continue
		}

		j, z := i+1, len(nums)-1
		for j < z {
			// 优化1
			if i+2 <= j && nums[j] == nums[j-1] {
				j++
				continue
			}

			if nums[i]+nums[j]+nums[z] < 0 {
				j++
			} else if 0 < nums[i]+nums[j]+nums[z] {
				z--
			} else {
				response = append(response, []int{nums[i], nums[j], nums[z]})
				j++
			}
		}
	}
	return response
}

func threeSum(nums []int) [][]int {
	// FAST PATH
	if len(nums) < 3 {
		return [][]int{}
	}
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		}
		return [][]int{}
	}

	sort.Ints(nums)
	var response [][]int
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		j, z := i+1, len(nums)-1
		for j < z {
			item := []int{nums[i], nums[j], nums[z]}
			sum := sum(item)
			if sum < 0 {
				j++
			} else if 0 < sum {
				z--
			} else {
				if !exist(response, item) {
					response = append(response, item)
				}
				j++
				z--
			}
		}
	}
	return response
}

func exist(l [][]int, arr []int) bool {
	for _, v := range l {
		if v[0] == arr[0] && v[1] == arr[1] && v[2] == arr[2] {
			return true
		}
	}
	return false
}

func sum(l []int) int {
	return l[0] + l[1] + l[2]
}
