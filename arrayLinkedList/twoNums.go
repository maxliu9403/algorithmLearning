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

func twoSun(nums []int, target int) [2]int {
	var resIndex [2]int

	for i := 0; i < len(nums); i++ {

		for j := i + 2; j < len(nums); j++ {
			value := nums[i] + nums[j]
			if value == target {
				resIndex[0] = i
				resIndex[1] = j
			}
		}
	}
	return resIndex

}

/*
思路2：两遍哈希法
1.循环创建一个包含所有元素的hash表，key是元素，value是元素下标
2.计算出目标值和元素的差值
3.循环判断差值是否存在再hash表中，如果存在且此时的下标不等于当前遍历的下标（保证不是重复元素）
*/

func twoSum2(nums []int, target int) [2]int {
	var resIndex [2]int
	Temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		Temp[nums[i]] = i
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

/*
思路：一遍hash
时间复杂度：O(n)
我们只遍历了包含有 n 个元素的列表一次。在表中进行的每次查找只花费 O(1)O(1) 的时间。

空间复杂度：O(n)
所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 n 个元素。
*/

func twoSum2(nums []int, target int) []int {
	Temp := make(map[int]int)
	//var resIndex [2]int
	resIndex := make([]int, 2, 10)
	for i := 0; i < len(nums); i++ {

		// 计算差值
		complement := target - nums[i]

		// 判断差值是否存在hash中
		// 如果存在则表示存在两数之和等于目标值
		// 如果不存在则将差值添加到hash表中
		index, ok := Temp[complement]
		if ok {
			resIndex[0] = i
			resIndex[1] = index
		} else {
			Temp[nums[i]] = i
		}

	}
	return resIndex
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 18
	resIndex := twoSum2(nums, target)
	fmt.Println(resIndex)
}
