package twoSum

import "sort"

// 时间复杂度O(n)
// 空间复杂度O(n)
func TwoSum(nums []int, target int) []int {
	ans := make([]int, 0)
	tmp := make(map[int]int, 0)
	for k, v := range nums {
		tmp[v] = k
	}
	for k, v := range nums {
		gap := target - v
		index, ok := tmp[gap]
		if ok && index != k {
			ans = []int{k, index}
			return ans
		}
	}
	return ans
}

func Method2(nums []int, target int) []int {
	tmp := make(map[int]int)
	for k, v := range nums {
		if index, ok := tmp[target-v]; ok {
			return []int{k, index}
		}
		tmp[v] = k
	}
	return nil
}

// 两数之和变化，输入nums和一个目标target，请返回nums中能够凑出target的两个元素，比如输入nums = [1, 3, 5, 6], target=9，而且接口集不重复,输出[3, 6]
func twoSumTarget(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			ans = append(ans, []int{nums[left], nums[right]})
			left++
			// 去掉重复值
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			right--
			// 去掉重复值
			for left < right && nums[right] == nums[right+1] {
				right--
			}

		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return ans
}
