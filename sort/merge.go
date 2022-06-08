package sort

import "sort"

/*
合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] <= intervals[j][1]
	})
	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res) - 1][1] < interval[0] {
			res = append(res, interval)
		}else {
			res[len(res) - 1][1] = max(res[len(res) - 1][1], interval[1])
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
时间复杂度：O(nlogn)
*/
