package twoSum

// 时间复杂度O(n)
// 空间复杂度O(n)
func TwoSum(nums []int, target int) []int {
	ans := make([]int, 0)
	tmp := make(map[int]int, 0)
	for k, v := range nums {
		tmp[v] = k
	}
	for k, v := range nums {
		gap := target - v
		index, ok := tmp[gap]
		if ok && index != k {
			ans = []int{k, index}
			return ans
		}
	}
	return ans
}
