package removeDuplicates

/*
	思路：双指针
*/
func removeDuplicatesMaxLiu(nums []int) int {
	n := len(nums)
	if n == 0 || nums == nil {
		return 0
	}
	left := 0
	right := 1
	for right < n {
		// 因为是有序递增数组，相同元素只能是相邻元素
		// 如果相邻元素不相等，将左指针右移到此时的右指针位置，同时右指针也右移一位
		if nums[left] != nums[right] {
			// 优化:如果数组中没有重复元素，此时将就会将left指向的元素原地复制一遍，这个操作其实是不必要的
			if (right - left) > 1 {
				nums[left+1] = nums[right]
			}
			left++
		}
		right++
	}
	// 因为left是从0开始的
	return left + 1
}

// 复杂度分析：
// 时间复杂度：O(n)
// 空间复杂度：O(1)
