package merge2

func merge(nums1 []int, m int, nums2 []int, n int) {
	resp := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if m <= i {
			resp = append(resp, nums2[j:]...)
			break
		}
		if n <= j {
			resp = append(resp, nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			resp = append(resp, nums1[i])
			i++
		} else {
			resp = append(resp, nums2[j])
			j++
		}
	}
	copy(nums1, resp)
}
