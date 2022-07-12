package intersect

import "sort"

/*
	思路：排序 + 双指针
	时间复杂度：O(mlogm+ nlogn),对两个数组排序的时间复杂度分别是O(mlogm)和O(nlogn)，双指针寻找交集元素的时间复杂度是O(m+n)，因此总时间复杂度是O(mlogm+nlogn)。

	空间复杂度：O(logm+logn)，空间复杂度主要取决于排序使用的额外空间
*/
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			ans = append(ans, x)
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return ans
}

/*
 	解法2；hash表
	时间复杂度：O(m+n)
	空间复杂度：O(min(m, n))
*/
func intersectMethod2(nums1 []int, nums2 []int) []int {
	var ans []int
	// 较短的数组并在哈希表中记录每个数字以及对应出现的次数
	if len(nums1) > len(nums2) {
		return intersectMethod2(nums2, nums1)
	}
	// 计数
	tmp := map[int]int{}
	for _, num := range nums1 {
		tmp[num] ++
	}
	for _, num := range nums2 {
		// 保证插入的交集元素是最少的
		if tmp[num] > 0 {
			ans = append(ans, num)
			tmp[num]--
		}
	}
	return ans
}
