package main

import (
	"fmt"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。
链接：https://leetcode-cn.com/problems/3sum
*/

/*
计算三数之和，其实就是判断a+b=-c的问题
*/

/*
思路1：暴力解法，三层循环


1.第一层循环计算出nums，所有的负数值
2.
2.计算出此时的离目标值的差值
3.二层循环利用夹逼发计算

*/
func threeSum(nums []int) [][]int {
	//
}

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	resIndex := threeSum(nums)
	fmt.Println(resIndex)
}
