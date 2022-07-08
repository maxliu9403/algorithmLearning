package missingNumber

import "sort"

/*
	排序：时间复杂度：O(nlogn)；空间复杂度：O(nlogn)
*/
func missingNumberSort(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}
	return len(nums)
}

/*
	数学：时间复杂度：O(n)；空间复杂度：O(1)
*/
func missingNumberMath(nums []int) int {
	n := len(nums)
	total := (n + 1) * n / 2
	a := 0
	for _, num := range nums {
		a += num
	}
	return total - a
}

/*
	排序 + 二分查找
*/
func missingNumberMaxLiu(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		// 索引和值相等，由于是排序的所以[left, mid]不存在错位，更新left的位置
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
