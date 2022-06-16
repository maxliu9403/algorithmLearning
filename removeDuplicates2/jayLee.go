package removeDuplicates2

// 哈希表解法
func removeDuplicatesHashTable(nums []int) int {
	h := make(map[int]int, len(nums))
	j := 0
	for _, item := range nums {
		dup, ok := h[item]
		if !ok || dup <= 1 {
			h[item] += 1
			nums[j] = item
			j++
		}
	}
	return len(nums[:j])
}

// 双指针解法
// 有序 & 最多重复2个 => 数组的前2位必然可以保留
func removeDuplicatesTwoPointers(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	low, fast := 2, 2
	for fast < len(nums) {
		// 如果fast值与low-2值相同，说明已经出现超过2个重复了，继续右移fast指针，直到出现不重复的元素为止。
		// 出现不重复的元素，low，fast指针都可以右移一位了。
		if nums[low-2] != nums[fast] {
			nums[low] = nums[fast] // 将出现的不重复的值覆盖到low所在的位置。
			low++
		}
		fast++
	}
	return low
}
