package merge

import (
	"math"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resp := make([][]int, len(intervals))
	idx := -1
	for _, interval := range intervals {
		// resp空数组 || 左区间 大于 resp最后一位的右区间 = 不重叠。
		if idx == -1 || interval[0] > resp[idx][1] {
			idx++
			resp[idx] = interval
		} else {
			resp[idx][1] = int(math.Max(float64(resp[idx][1]), float64(interval[1])))
		}
	}
	return resp[:idx+1]
}
