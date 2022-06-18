package merge

import "sort"

/*
	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
	思路：自定义排序
	时间复杂度：自定义排序O(nlogn) + 遍历O(n) = O(nlogn) + O(n) = O(nlogn)
	空间复杂度：O(logn) ，其中n为区间的数量
*/
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0, len(intervals))
	// 自定义排序, 保证整个intervals是按照子切片的0号元素的有序递增
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})
	for _, interval := range intervals {
		// len(ans) == 0 遍历第一个元素直接加入结果集
		// ans[len(ans)-1][1] < interval[0] 结果集中的最后一个元素的右边界 < 新元素的左边界，则不存在重复区间
		if len(ans) == 0 || ans[len(ans)-1][1] < interval[0] {
			ans = append(ans, interval)
			// 更新结果集中右边界
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/*
	思路：自定义排序 + 双指针
	时间复杂度：自定义排序O(nlogn) + 遍历O(n) = O(nlogn) + O(n) = O(nlogn)
	空间复杂度：O(logn) ，其中n为区间的数量
*/
func merge2(intervals [][]int) [][]int {
	ans := make([][]int, 0, len(intervals))
	// 自定义排序,保证整个intervals是一个有序递增
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})
	n := len(ans)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] > right {
			ans = append(ans, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}else {
			right = max(right, intervals[i][1])
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
