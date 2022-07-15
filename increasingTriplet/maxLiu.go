package increasingTriplet

import "math"

/*
	思路：双向遍历
	下标：i < j < k，存在nums[i] < nums[j] < nums[k]
	假设存在下标i, 1 <= i < n-1 存在i左边的一个元素使得<nums[i]，存在i右边的一个元素使得>nums[i]
    创建两个长度为n，leftMin,rightMax,对于0<=i<n, leftMin[i] 表示nums[0] 到 nums[i] 中的最小值，rightMax[i] 表示 nums[i] 到 nums[n−1] 中的最大值
*/
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		// 在i的左边寻找比nums[i]小的元素
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 贪心解法
赋初始值的时候，已经满足second > first了，现在找第三个数third
(1) 如果third比second大，那就是找到了，直接返回true
(2) 如果third比second小，但是比first大，那就把second的值设为third，然后继续遍历找third
(3) 如果third比first还小，那就把first的值设为third，然后继续遍历找third（这样的话first会跑到second的后边，但是不要紧，因为在second的前边，老first还是满足的）
*/
func increasingTripletMethod2(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		third := nums[i]
		if third > second {
			return true
		} else if third > first {
			second = third
		} else {
			first = third
		}
	}
	return false
}
