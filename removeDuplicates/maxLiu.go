package removeDuplicates

/*
	思路：双指针
*/

// 题1：给你一个 升序排列 的数组 nums ，请你原地删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么nums的前 k 个元素应该保存最终结果

// 复杂度分析：
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicatesMaxLiu(nums []int) int {
	n := len(nums)
	if n == 0 || nums == nil {
		return 0
	}
	left := 1
	right := 1
	for right < n {
		// 因为是有序递增数组，相同元素只能是相邻元素
		// 如果相邻元素不相等，将左指针右移到此时的右指针位置，同时右指针也右移一位
		if nums[left-1] != nums[right] {
			// 优化:如果数组中没有重复元素，此时将就会将left指向的元素原地复制一遍，这个操作其实是不必要的
			if (right - left) > 0 {
				nums[left] = nums[right]
			}
			left++
		}
		right++
	}
	// 因为left是从0开始的
	return left
}
