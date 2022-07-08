package hIndex

import "sort"

/*
	题意：某人的 h 指数是 20，这表示他已发表的论文中，每篇被引用了至少 20 次的论文总共有 20 篇，所以有h<=20
	思路：排序；令h=0，从后向前遍历数组，判断篇论文应用次数是否大于h，如果大于h继续加一，直到当前论文引用次数小于h
	复杂度：时间复杂度：O(nlogn); 空间O(logn)
*/
func hIndexMaxLiu(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i]>h{
			h++
		}
	}
	return h
}