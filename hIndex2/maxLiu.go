package hIndex2

/*
 题意：某人的 h 指数是 20，这表示他已发表的论文中，每篇被引用了至少 20 次的论文总共有 20 篇，所以有h<=20，要求对数时间复杂度
 思路：二分查找
*/
func hIndex(citations []int) int {
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		// 猜论文篇数
		mid := (left + right) / 2
		// 满足高引用的特点是：引用次数 >= 论文篇数
		// count 的含义是：大于等于 mid 的论文篇数
		count := 0
		for _, citation := range citations {
			if citation >= mid {
				count++
			}
		}
		// 如果大于等于 mid 的论文篇数大于等于 mid，说明 h 指数至少是 mid
		if count >= mid {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
