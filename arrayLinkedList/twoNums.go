package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那
两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
链接：https://leetcode-cn.com/problems/two-sum
*/

/*
思路1：暴力解法（两层for）
对于每个元素，我们试图通过遍历数组的其余部分来寻找它所对应的目标元素，这将耗费 O(n)O(n) 的时间。
因此时间复杂度为 O(n^2)
空间复杂度：O(1)O(1)。

*/

func twoSum(nums []int, target int) [2]int {
	var resIndex [2]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			value := nums[i] + nums[j]
			if value == target {
				resIndex[0] = i
				resIndex[1] = j
			}
			continue
		}
	}

	return resIndex

}

/*
思路2：两遍哈希法

*/

func twoSum1(nums []int, target int) [2]int {

	var resIndex [2]int
	Temp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		Temp[nums[i]] = i
		/*
		{2:0, 7:1, 11:2, 15:3}
		*/
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		_, ok := Temp[complement]
		// 数组中同一个元素不能使用两遍。
		if ok && Temp[complement] != i {
			resIndex[0] = i
			resIndex[1] = Temp[complement]

		}

	}
	return resIndex

}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 18
	resIndex := twoSum1(nums, target)
	fmt.Println(resIndex)
}
