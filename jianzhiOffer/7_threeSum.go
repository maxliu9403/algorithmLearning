package jianzhiOffer

import "sort"

func threeSum(nums []int, target int) [][]int {
	return nSumTarget(nums, 3, 0, target)
}

func nSumTarget(nums []int, n int, start, target int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// 至少是两数之后，同时数组长度必须大于计算数字之和的元素个数
	sz := len(nums)
	if n < 2 || sz < n {
		return res
	}
	// base case
	if n == 2 {
		lo := start
		hi := sz - 1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				// 过滤重复元素
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				// 过滤重复元素
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{left, right})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// n>2时递归计算
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}
	return res
}
