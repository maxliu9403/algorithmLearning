package wiggleSort

import (
	"sort"
)

// todo 快选 + 三数切分
func wiggleSort(nums []int) {
	numsBak := append([]int{}, nums...)
	sort.Ints(numsBak)
	n := len(nums)
	// 将数组分为2部分，[(n+1)/2]
	for idx, i, j := 0, (n+1)/2-1, n-1; idx < n; idx++ {
		if idx%2 == 0 { // i
			nums[idx] = numsBak[i]
			i--
		} else {
			nums[idx] = numsBak[j]
			j--
		}
	}
}
