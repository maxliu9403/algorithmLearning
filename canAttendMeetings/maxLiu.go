package canAttendMeetings

import "sort"

func canAttendMeetingsMaxLiu(intervals [][]int) bool {
	// 按照会议的开始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// [[0,30],[5,10],[15,20]]
	// 前一个会议的结束时间大于当前会议的开始时间
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
