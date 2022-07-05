package containsNearbyAlmostDuplicate

import (
	"math/rand"
	"time"
)

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	rand.Seed(time.Now().UnixNano())
	QuickSort(nums, 0, len(nums)-1)
	j, i := 0, 1
	for i <= len(nums)-1 {
		if absFunc(i, j) <= k && absFunc(nums[i], nums[j]) <= t {
			return true
		}
		i++
		j++
	}
	return false
}

func absFunc(i, j int) int {
	if i-j < 0 {
		return j - i
	}
	return i - j
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
