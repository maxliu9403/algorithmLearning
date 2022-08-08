package jianzhiOffer

func numSubarrayProductLessThanK(nums []int, k int) int {
	// 初始值为1
	left, product := 0, 1
	count := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		product *= nums[right]
		for left <= right && product >= k {
			// 移动左边界
			product = product / nums[left]
			left++
		}
		// 乘积小于k，那么数组中每个元素都是小于k的
		count += right - left + 1
	}
	return count
}
