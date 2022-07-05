package missingNumber

import (
	"sort"
)

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

/*
Thinking:
	1. 异或:n^n=0, n^0=n
	2. 将数组扩充为:[0..len(nums)] + nums, 这样就确保nums内的所有元素都能被抵消为0,
	   最终会得到0^missingNumber = missingNumber
*/
func missingNumberXor(nums []int) int {
	var xor int
	for i, num := range nums {
		// 这是一种缩略的写法，使省略了生成[0..len(nums)-1]数组的时间和空间
		// 但是，会少异或一个元素：len(nums)
		xor ^= num ^ i
	}
	return xor ^ len(nums) // 异或len(nums)
}
