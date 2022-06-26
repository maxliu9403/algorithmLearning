package largestNumber

import (
	"sort"
	"strconv"
)

// [111311, 1113]
func largestNumber(nums []int) string {
	var sortF = func(i, j int) bool {
		// 比较111311,1113 是否 大于 1113,111311
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		// sx: x的进位数
		for sx <= x {
			sx *= 10
		}
		// sy: y的进位数
		for sy <= y {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	}
	// 按高低位降序排序
	sort.Slice(nums, sortF)
	// 降序排序后首位为0的情况
	if nums[0] == 0 {
		return "0"
	}
	// 将排序后的数组拼接
	var ans []byte
	for _, num := range nums {
		ans = append(ans, strconv.Itoa(num)...)
	}
	return string(ans)
}
