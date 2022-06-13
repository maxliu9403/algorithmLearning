package maxArea

func maxArea(height []int) int {
	left, right, maxS := 0, len(height)-1, 0
	for left < right {
		if s := (right - left) * min(height[left], height[right]); maxS < s {
			maxS = s
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxS
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
