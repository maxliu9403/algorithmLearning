package twoSum

func twoSum(nums []int, target int) []int {
	h := map[int]int{}
	for i := range nums {
		if _, ok := h[target-nums[i]]; ok {
			return []int{h[target-nums[i]], i}
		}
		h[nums[i]] = i
	}
	return []int{}
}
