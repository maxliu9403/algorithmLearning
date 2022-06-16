package removeElement

func removeElement(nums []int, val int) int {
	idx := 0
	for _, item := range nums {
		if item != val {
			nums[idx] = item
			idx++
		}
	}
	return len(nums[:idx])
}
