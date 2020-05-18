package main

import (
	"fmt"
	"math"
)

// 【盛最多的水】11题

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
链接：https://leetcode-cn.com/problems/container-with-most-water
*/

/*
思路1：枚举法
计算出每种区域的面积，选择出最大值
时间复杂度：O(n^2)

*/

func maxArea1(height []int) float64 {
	var maxValue float64
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
			maxValue = math.Max(float64(area), maxValue)
		}
	}
	return maxValue
}

// golang对于int/int64类型的数据，不提供直接的min和max操作
func maxArea3(height []int) int {
	var maxValue int
	var resValue int
	for i := 0; i < len(height)-1; i++ {
		for j := 0; j < len(height); j++ {

			if height[i] < height[j] {
				resValue = (j - i) * height[i]
			} else {
				resValue = (j - i) * height[j]
			}

			// 选择出最大值
			if maxValue <= resValue {
				maxValue = resValue
			}
			continue

		}
	}
	return maxValue

}

/*
思路2：夹逼法
1.双指针
2.初始指针分别放再0第一个元素和最后一个元素
3.移动左边指针，如果此时的指向的左边元素的高小于上次的左边元素，
跳过（因为此时的宽度变窄了，高度也变小了，那么自然面积变小了）

时间复杂度：O(N)
*/

// 枚举法
func maxArea2(height []int) int {
	size := len(height)
	left, right := 0, size-1
	var maxValue int

	for left != right {
		var resValue int
		if height[left] < height[right] {
			resValue = height[left] * (right - left)
			left++
		} else {
			resValue = height[right] * (right - left)
			right--
		}
		// 比较最大值
		if maxValue >= resValue {
			continue
		} else {
			maxValue = resValue
		}
	}
	return maxValue

}


// 夹逼法
func maxArea(nums []int) int {
	left, right := 0, len(nums)-1
	var maxValue int

	for left != right {
		var resValue int
		if nums[left] < nums[right] {
			resValue = nums[left] * (right - left)
			// 右移动
			left++
		} else {
			resValue = nums[right] * (right - left)
			// 左移
			right--
		}

		// 选择出最大值
		if maxValue >= resValue {
			continue
		} else {
			maxValue = resValue
		}
		continue

	}
	return maxValue

}

func main() {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxValue := maxArea(height)
	fmt.Println(maxValue)

}
