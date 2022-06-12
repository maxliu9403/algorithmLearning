package threeSumClosest

import (
	"sort"
)

/**
题目：
	给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
	返回这三个数的和。
	假定每组输入只存在恰好一个解。

思路：
	遍历从小到大排序后的数组，放置双指针的位置为：下一位和最后一位，求出3个位置的和 sum。
	计算 sum 与 target 的绝对距离，如果比上次距离更近，则替换之。
	1. 如果 sum 正好等于 target：由于只存在恰好一个解，判断为最近，直接返回sum。
	2. 如果 sum 小于 target：说明整体位置偏左，左指针右移。
	3. 如果 sum 大于 target：说明整体位置偏右，右指针左移。
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	response := nums[0] + nums[1] + nums[2]
	dev := abs(response - target)
	for i := 0; i < len(nums); i++ {
		j, z := i+1, len(nums)-1
		for j < z {
			sum := nums[i] + nums[j] + nums[z]
			if dev1 := abs(sum - target); dev1 < dev {
				dev = dev1
				response = sum
			}
			if sum < target {
				j++
			} else if target < sum {
				z--
			} else {
				return sum
			}
		}
	}
	return response
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
