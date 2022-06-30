package containsDuplicate

func containsDuplicate(nums []int) bool {
	h := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := h[num]; !ok {
			h[num] = struct{}{}
			continue
		}
		return true
	}
	return false
}
