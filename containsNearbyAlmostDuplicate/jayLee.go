package containsNearbyAlmostDuplicate

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// 因为sort使用的是原址排序，因此无法得到原下标。
	// 本题适合用快慢指针解题。确保 f - l <= k 即 f <= k + l
	length := len(nums)
	for l := 0; l < length; l++ {
		f := l + k // f
		if length-1 < f {
			f = length - 1
		}
		for f > l {
			if abs(nums[l]-nums[f]) <= t {
				return true
			}
			f--
		}
	}
	return false
}

func abs(s int) int {
	if s < 0 {
		return -s
	}
	return s
}
