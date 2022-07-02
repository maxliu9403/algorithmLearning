package majorityElement

// majorityElementBoyerMoore 摩尔选举法。相似题目：majorityElement2.MajorityElement
// nums:       [7, 7, 5, 7, 5, 1 | 5, 7 | 5, 5, 7, 7 | 7, 7, 7, 7]
// candidate:  [7, 7, 7, 7, 7, 7,  5, 5,  5, 5, 5, 5,  7, 7, 7, 7]
// count:      [1, 2, 1, 2, 1, 0,  1, 0,  1, 2, 1, 0,  1, 2, 3, 4]
func majorityElementBoyerMoore(nums []int) int {
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
		} else if count > 0 && candidate != num {
			count--
		} else if count > 0 && candidate == num {
			count++
		}
	}
	return candidate
}

// 找出第一个超过 nums.length / 2 的数字
func majorityElementHashTable(nums []int) int {
	h := map[int]int{} // 由于用到了hash table所以空间复杂度最大为O(n)
	t := len(nums) / 2
	for _, num := range nums {
		h[num]++
		if h[num] > t {
			return num
		}
	}
	return 0
}

// TODO 分治法
func majorityElementDivideAndConquer(nums []int) int {
	var sum = func(t int) (count int) {
		for _, n := range nums {
			if n == t {
				count++
			}
		}
		return
	}

	var recursion func(l, h int) int
	recursion = func(l, h int) int {
		// case1：如果只有1个数，一定为众数，直接返回。
		if l == h {
			return nums[l]
		}

		// case2：如果为多个数，左右拆分为2个子问题，分别递归。
		mid := (h-l)/2 + l
		left := recursion(l, mid)
		right := recursion(mid+1, h)

		leftCount := sum(left)   // 左边的众数的个数
		rightCount := sum(right) // 右边众数的个数

		if leftCount > rightCount {
			return left
		} else {
			return right
		}
	}
	return recursion(0, len(nums)-1)
}
