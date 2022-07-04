package findKthLargest

import (
	"math/rand"
	"sort"
	"time"
)

// 解法1：暴力解法排序
// 寻找第k大的元素
// [1, 2, 3, 4, 5, 6] k=1,index = 5, k=3,index=3 index = len(nums)-k
func findKthLargestMethod1(nums []int, k int) int {
	var ans int
	sort.Ints(nums)
	n := len(nums)
	ans = nums[n-k]
	return ans
}

// 解法2：减而治之（逐渐缩小问题规模）
func findKthLargestPartition(nums []int, k int) int {
	n := len(nums)
	targetIndex := n - k
	// 在区间nums[left, right]区间内寻找target
	left := 0
	right := n - 1
	for true {
		pivotIndex := partitionFunc(nums, left, right)
		if pivotIndex == targetIndex {
			return nums[pivotIndex]
		} else if pivotIndex < targetIndex {
			left = pivotIndex + 1
		} else if pivotIndex > targetIndex {
			right = pivotIndex - 1
		}
	}
	return nums[targetIndex]
}

// 解法三：减而治之 + 快速排序
func findKthLargestQuickSort(nums []int, k int) int {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	targetIndex := n - k
	// 在区间nums[left, right]区间内寻找target
	left, right := 0, n-1
	QuickSort(nums, left, right)
	return nums[targetIndex]
}

// 快速排序
func QuickSort(nums []int, left, right int) {
	// 终止条件是此时数组中没有元素
	if left >= right {
		return
	}
	// 确定划分位置
	pivotIndex := partitionFunc(nums, left, right)
	// 递归排序
	// 对分区左边元素进行排序
	QuickSort(nums, left, pivotIndex-1)
	// 对分区右边元素进行排序
	QuickSort(nums, pivotIndex+1, right)
}

/*
分区：将一个数组划分为三个部分
[nums1 , pivot, nums2]
all num in nums1 <= pivot <  all num in nums2
*/
func partitionFunc(nums []int, left, right int) int {
	// 定义分区比较元素
	// 随机数快速排序
	//
	/*
		因为partition的顺序逆序排序中表现很差
		拆分的子问题的规模只比原来减少一个元素，每次排序只能确定一个元素
	*/
	randomIndex := rand.Intn(right-left) + left
	swap(nums, left, randomIndex)
	pivot := nums[left]
	// [left+1, j] <= pivot
	// (j, i) > pivot
	// j是划分第一个区间时的最后一个元素
	j := left
	for i := left + 1; i <= right; i++ {
		// 当前元素小于等于分区比较元素，因为j是划分第一个区间时的最后一个元素，所以j+1，和i交换，保证交换后i所在的位置是第一个区间的最后一个元素
		if nums[i] <= pivot {
			j++
			swap(nums, i, j)
		}
	}
	// 将分区元素和第一个区间的最后一个元素进行交换
	swap(nums, left, j)
	return j
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
