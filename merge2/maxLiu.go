package merge2

import "sort"

// 直接排序
// 将数组nums2放入nums1中后，然后进行排序
/*
	时间复杂度：O((m+n)log(m+n))
排序序列长度为 m+nm+n，套用快速排序的时间复杂度即可，平均情况为O((m+n)log(m+n))

	空间复杂度：O(log(m+n))。
排序序列长度为 m+nm+n，套用快速排序的空间复杂度即可，平均情况为O(log(m+n))。
*/
func merge1(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

/*
	思路：
	双指针，两个数组看作队列从头部取出比较小的数字放到结果中

*/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	// p1 表示nums1的指针，初始值在0号位置
	// p2 表示nums2的指针，初始值在0号位置
	// 分别遍历nums1和nums2
	p1, p2 := 0, 0
	for {
		// 必须先判断退出条件，否则下面进行条件判断时会出现超出索引的问题
		// 表示此时nums1已经全部遍历完毕了，将num2中剩下的元素全部添加即可
		if p1 == m {
			// 此时的p2的取值范围0<=p2<n,不可能出现超出索引的问题
			// 因为如果存在p2>n，就必须有p2=n的情况
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		// 表示此时nums2已经全部遍历完毕了，将num1中剩下的元素全部添加即可
		if p2 == n {
			// 此时的p1的取值范围0<=p1<n
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		// 比较两个数组中指针所指向的元素，将小的元素放入，同时向后移动指针
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
