package removeElement

/*
	思路：双指针，通过判断左右指针元素与val的大小来进行指针的移动
*/
func removeElementMaxLiu(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 时间复杂度：O(n)O(n)
// 空间复杂度：O(1)O(1)
