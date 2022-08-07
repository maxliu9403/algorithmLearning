package threeSum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return ans
	}
	// 排序递增
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	// 第一层循环，选出1号元素
	for i := 0; i < n; i++ {
		// 因为已经排序好，所以后面不可能有三个数加和等于0，直接返回结果
		if nums[i] > 0 {
			return ans
		}
		// 去掉相同元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 定义左右指针
		l, r := i+1, n-1
		// 第二层循环，选出2号元素
		for l < r {
			// i <= l < r  =>  nums[i] <= nums[l] < nums[r]
			s := nums[i] + nums[l] + nums[r]
			// 三个元素==0，保持其中两个元素不变，改变3号元素
			if 0 == s {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				// 开始第3层循环，同时对3号元素进行去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 对2号元素进行去重
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// 移动2号元素
				l++
				r--
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

// 复杂度分析
// 时间复杂度：O(n^2) 数组排序O(NlogN)，遍历数组O(n)，双指针遍历O(n) 总体:O(NlogN) + O(n)*O(n)
// 空间复杂度：O(1)

func SumMe(nums []int) [][]int {
	return nSumTarget(nums, 4, 0, -6401)
}

// n数字之和
// n:几数之和
// start:枚举
// target:目标值
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
