package containsDuplicate

import (
	"math/rand"
	"time"
)

//
/*
	思路：排序后判断相邻元素是否相等
	随机数 + 快排 + 双指针
*/
func containsDuplicateMaxLiu(nums []int) bool {
	rand.Seed(time.Now().UnixNano())
	QuickSort(nums, 0, len(nums)-1)
	left, right := 0, 1
	for right <= len(nums)-1 {
		if nums[left] == nums[right] {
			return true
		}
		left++
		right++
	}
	return false
}

func partition(nums []int, left, right int) int {
	// 通过随机数保证left是随机的
	randomIndex := rand.Intn(right-left) + left
	swap(nums, left, randomIndex)
	pivot := nums[left]
	// [left+1, j] <= pivot
	// (j, i) > pivot
	j := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			j++
			swap(nums, i, j)
		}
	}
	swap(nums, left, j)
	return j
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(nums, left, right)
	QuickSort(nums, left, pivotIndex-1)
	QuickSort(nums, pivotIndex+1, right)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
