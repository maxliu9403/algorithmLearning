package findMedianSortedArrays

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n, idx := len(nums1), len(nums2), 0
	sortNums := make([]int, m+n)
	if n <= m {
		for i := 0; i < m; i++ {
			if n <= i {
				sortNums[idx] = nums1[i]
			} else {
				sortNums[idx] = nums2[i]
				sortNums[idx+1] = nums1[i]
				idx++
			}
			idx++
		}
	} else {
		for i := 0; i < n; i++ {
			if m <= i {
				sortNums[idx] = nums2[i]
			} else {
				sortNums[idx] = nums2[i]
				sortNums[idx+1] = nums1[i]
				idx++
			}
			idx++
		}
	}
	sort.Ints(sortNums)
	// 奇数的中位数索引 = (m+n) / 2
	if (m+n)%2 == 1 {
		return float64(sortNums[(m+n)/2])
	}
	// 偶数数的中位数索引 = [((m+n) / 2) -1, (m+n) / 2]
	return float64(sortNums[((m+n)/2)-1]+sortNums[(m+n)/2]) / 2
}

// @lc code=end
