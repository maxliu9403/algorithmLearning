package intersection

import "sort"

/*
	时间复杂度：O(n+n),空间复杂度:O(m+n)
*/
func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	tmp := map[int]struct{}{}
	for _, num := range nums1 {
		tmp[num] = struct{}{}
	}
	filter := map[int]int{}
	for _, num := range nums2 {
		_, ok := tmp[num]
		filter[num] ++
		if ok && filter[num] <= 1 {
			ans = append(ans, num)
		}
	}
	return ans
}

/*
	思路：排序 + 双指针
	时间复杂度：O(mlogm+ nlogn),对两个数组排序的时间复杂度分别是O(mlogm)和O(nlogn)，双指针寻找交集元素的时间复杂度是O(m+n)，因此总时间复杂度是O(mlogm+nlogn)。

	空间复杂度：O(logm+logn)，空间复杂度主要取决于排序使用的额外空间
*/
func intersectionMethod2(nums1 []int, nums2 []int) []int {
	var ans []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	var pre int // 记录前一个交集元素
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		// 分别取出两个数组的元素
		x, y := nums1[i], nums2[j]
		// 交集元素
		if x == y {
			if len(ans) == 0 || pre != x {
				ans = append(ans, x)
				// 更新前一个交集元素
				pre = x
			}
			// 同时移动两个指针
			i++
			j++
		} else if x < y {
			// 移动较小元素的下标
			i++
		} else {
			j++
		}
	}
	return ans
}
