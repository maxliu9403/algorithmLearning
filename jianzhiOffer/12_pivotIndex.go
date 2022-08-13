package jianzhiOffer

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		v := nums[i]
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
