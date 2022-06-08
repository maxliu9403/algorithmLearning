package sort

import (
	"sort"
)

/*
数组相对排序
给定两个数组，arr1和arr2，
arr2中的元素各不相同arr2中的每个元素都出现在arr1中对arr1中的元素进行排序，使arr1中项的相对顺和arr2中的相对顺序相同。未在arr2中出现过的元素需要按照升序放在arr1的末尾
*/

// 方法1：自定排序
func RelativeSortArrayMethod1(arr1 []int, arr2 []int) []int {
	// 由于已知arr2的排序规则，通过map记录这个规则
	// 比如：arr2 = [3, 7, 1] rank = {3:0, 7:1, 1:2}
	// 当arr1中存在arr2中元素时通过元素值取出下标则是排序规则
	rank := make(map[int]int)
	for i, v := range arr2 {
		rank[v] = i
	}
	// 自定排序，返回true则保持原来的排序，返回false则交换位置
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		// x元素在arr2中位置
		rankx, hasx := rank[x]
		// y元素在arr2中位置
		ranky, hasy := rank[y]
		// 如果同时存在x和y，则通过rank的值进行排序比较
		if hasx && hasy {
			return rankx < ranky
		}
		// 题目要求未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾，
		// 因为i>j, x是排在y的后面，所以如果map中存在x则此时x与y需要交换顺序
		if hasx || hasy {
			return hasx
		}
		// 如果x和y都没有出现在哈希表中，那么比较它们本身
		return x < y
	})
	return arr1
}

// 方法2：计数排序
func RelativeSortArrayMethod2(arr1 []int, arr2 []int) []int {
	// 找出arr1中的最大值
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	// 记录arr1各个元素出现的频率, 数组的长度最大为arr1中的最大元素+1
	// 索引是arr1 中的元素，值为 arr1 中的元素出现的次数
	// 此时的索引其实就是arr1不在arr2元素的递增排序了
	frequency := make([]int, upper+1)
	//arr1 := []int{19, 2, 3, 1, 3, 2, 4, 6, 7, 9, 2}
	//arr2 := []int{2, 1, 4, 3, 9, 6}
	for _, v := range arr1 {
		frequency[v] ++
	}
	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	// 处理arr1中不在arr2中的元素
	// v：arr1元素
	// freq：v出现的次数
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}

/*
计数排序：
时间复杂度：O(m+n+upper)
空间复杂度：O(upper)，即为数组frequency需要使用的空间

*/
