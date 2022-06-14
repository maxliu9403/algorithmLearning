package removeDuplicates

func removeDuplicates(nums []int) int {
	h := make(map[int]struct{}, len(nums))
	j := 0
	for _, item := range nums {
		_, ok := h[item]
		if !ok {
			h[item] = struct{}{}
			nums[j] = item
			j++
		}
	}
	return len(nums[:j])
}
