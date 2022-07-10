package wiggleSort

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)
	ans := make([]int, 0)
	left, right := (len(nums)-1)/2, len(nums)-1
	for left >= 0 {
		ans = append(ans, nums[left])
		ans = append(ans, nums[right])
		left--
		right--
	}
	copy(nums, ans)
}
