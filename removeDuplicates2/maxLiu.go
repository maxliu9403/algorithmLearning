package removeDuplicates2

// 给你一个有序数组nums ，请你原地删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

/*
	思路：双指针
	有序数组 相同元素相临
*/

func RemoveDuplicatesMedium(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	left := 2
	right := 2
	for right < n {
		// 因为是有序递增数组，相同元素只能是相邻元素
		// 如果相邻元素不相等，将左指针右移到此时的右指针位置，同时右指针也右移一位
		if nums[left-2] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// 一般情况：给你一个有序数组nums ，请你原地删除重复出现的元素，使每个元素 最多出现K
//次，返回删除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicatesK(nums []int, k int) int {
	n := len(nums)
	if n <= k {
		return n
	}
	left := k
	right := k
	for right < n {
		// 因为是有序递增数组，相同元素只能是相邻元素
		// 如果相邻元素不相等，将左指针右移到此时的右指针位置，同时右指针也右移一位
		if nums[left-k] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
