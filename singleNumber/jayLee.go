package singleNumber

// XorSingleNumber
/*
Thinking 异或运算
	1. 首先，找出2个唯一元素的异或值，xor = t1 ^ t2，可以保证的是：两个不同元素异或，至少有1位是非0的。
	例如：3 ^ 5 = 6
		011 = 3
	  ^	101 = 5
	 ------------
		110 = 6
	2. 找出其中异或值为1的位 h，一定有该特性：t1和t2一定有一个为1，一个为0。
	3. 将h与nums的元素，如果与值为0的放入一组，与值非1的放入另一组，依次异或。
Association: missingNumber.XorMissingNumber
*/
func XorSingleNumber(nums []int) []int {
	var t1, t2 int
	var xor int
	for _, num := range nums {
		xor ^= num
	}

	h := 1
	for xor&h == 0 {
		h <<= 1
	}
	for _, num := range nums {
		if h&num == 0 {
			t1 ^= num
		} else {
			t2 ^= num
		}
	}
	return []int{t1, t2}
}
