package largestNumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumberMaxLiu(nums []int) string {
	ss := make([]string, len(nums))
	// 转换为字符串
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	// 比较两个相邻元素组合后大小，来进行排序
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	ans := strings.Join(ss, "")
	// 首位为0
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
