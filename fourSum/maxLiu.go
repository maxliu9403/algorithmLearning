package fourSum

import "sort"

/*
	给你一个由n个整数组成的数组nums ，和一个目标值target 。请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：

	0 <= a, b, c, d< n
	a、b、c 和 d 互不相同
	nums[a] + nums[b] + nums[c] + nums[d] == target
	你可以按 任意顺序 返回答案
*/
func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	if nums == nil || len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		// 去掉重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 计算当前最小值，如果当前最小值都是大于target，因为是递增的后面的元素就不可能了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 计算当前的最大值，如果当前的最大值小于target，后面的元素还有机会
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// 开启第二层循环
		// j < n-2 是为了后面在剪枝的时避免索引溢出
		// j := i + 1 是为了保证第二层循环和第一层循环的元素不一致
		for j := i + 1; j < n-2; j++ {
			// 去掉重复元素
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 计算当前的最小值，最小值大于target直接退出循环
			if j > i+1 && nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 计算当前的最大值，最大值小于target，后续还有希望
			if j > i+1 && nums[j]+nums[i]+nums[n-2]+nums[n-1] < target {
				continue
			}
			// 进行双指针遍历
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					// 先移动指针，和它前面的left-1进行比较，若后+1,则和它后面的left+1进行比较
					left++
					// 去掉相同元素
					for left < right && nums[left] == nums[left-1] {
						left++
						continue
					}
					// 先移动指针，和它后面的right+1进行比较，若后-1,则和它后面的right-1进行比较
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
						continue
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}
