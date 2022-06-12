package maxArea

// 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器
func MaxArea(height []int) int {
	// [1,8,6,2,5,4,8,3,7]
	// 双指针定义左右边界
	i, j := 0, len(height)-1
	var ans int
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		ans = max(area, ans)
		// 移动指针
		// 若向内 移动短板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
		//若向内 移动长板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j])不变或变小，因此下个水槽的面积 一定变小 。
		if height[i] <= height[j] {
			// 左指针小，左指针右移
			i++
		} else {
			j--
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 复杂度分析
// 时间复杂度：O(N)，双指针总计最多遍历整个数组一次。
// 空间复杂度：O(1)，只需要额外的常数级别的空间。
