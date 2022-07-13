package intersection

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var resp []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] { // 相等
			// 并且不等于pre
			if resp == nil || nums1[i] != resp[len(resp)-1] {
				resp = append(resp, nums1[i])
			}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return resp
}
