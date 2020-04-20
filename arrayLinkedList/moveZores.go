package main

import "fmt"

// 移动0元素到数组最右侧:LeetCode

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

/*
思路：
双指针思路：
第一次遇到非零元素，将非零元素与数组nums[0]互换
第二次遇到非零元素，将非零元素nums[1]互换
*/

/*
时间复杂度：O(n)
空间复杂度：
*/

func moveZeroes(nums []int) []int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if 0 != nums[j] {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
	}
	return nums

}

func main() {
	var nums = []int{0, 1, 0, 3, 12}
	newNums := moveZeroes(nums)
	fmt.Println(newNums)
}
