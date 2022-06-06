package main

import (
	"fmt"
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

func maxArea(height []int) int {
	var maxValue int
	var resValue int

	for i := 0; i < len(height)-1; i++ {

		for j := i + 1; j < len(height); j++ {
			if height[i] < height[j] {
				resValue = (j - i) * height[i]
			} else {
				resValue = (j - i) * height[j]
			}

		}

		// 选择最大值
		if maxValue <= resValue {
			maxValue = resValue
		}

	}
	return maxValue

}

/*
思路2：夹逼法
1.双指针
2.初始指针分别放再0第一个元素和最后一个元素
3.移动左边指针，如果此时的指向的左边元素的高小于上次的左边元素，跳过（因为此时的宽度变窄了，高度也变小了，那么自然面积变小了）
时间复杂度：O(N)
*/

//思路2：夹逼法
func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	var maxArea int

	for left != right {
		var resValue int
		if height[left] < height[right] {
			resValue = height[left] * (right - left)
			// 此时left右移
			left++

		} else {
			resValue = height[right] * (right - left)
			// 此时right左移
			right--

		}

		// 选出最大值
		if maxArea <= resValue {
			maxArea = resValue
		}

	}
	return maxArea

}

func main() {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxValue := maxArea2(height)
	fmt.Println(maxValue)

}
