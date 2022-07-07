package minMeetingRooms

import (
	"sort"
)

/*
Question
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],…] (si < ei)，为避免会议冲突，同时要考虑充分利用会议室资源，
请你计算至少需要多少间会议室，才能满足这些会议安排。
*/

/*
	思路1：暴力解法，按照开始时间排序，遍历每个会议判断当前会议的开始时间和上个会议的结束时间的关系。
	如果当前会议开始时间小于上一个会议的结束时间，需要新增会议（记录一个新的结束时间）
	如果当前会议开始时间大于上一个会议的结束时间，无需新增会议

*/
func minMeetingRoomsMaxLiu(intervals [][]int) int {
	// 按照会议的开始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 记录会议室
	rooms := make([]int, 0, len(intervals))
	for _, interval := range intervals {
		var find bool
		// 遍历所有已记录的会议室
		for index, endTime := range rooms {
			// 找到满足结束时间早于当前会议开始时间的会议室，并更新会议室的时间表
			// i3 := [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}}
			// [1, 10] rooms = [10]
			// [2, 7] rooms = [10, 7]
			// [3, 19] rooms = [10, 12, 19]
			// [8 ,12] rooms = [10, 12, 19]
			// [10, 20] rooms = [20, 12, 19]
			// [11, 30] rooms = [20, 12, 19, 30]
			if endTime <= interval[0] {
				rooms[index] = interval[1]
				find = true
				break
			}
		}
		// 没有找到合适的会议，新增会议室
		if !find {
			rooms = append(rooms, interval[1])
		}
	}
	return len(rooms)
}
