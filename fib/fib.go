package main

import "fmt"

// 斐波那契数列
/*
暴力解法
1.时间复杂度O(2^n)
2.重叠子问题
*/

//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//
//	if n == 1 || n == 2 {
//		return 1
//	}
//
//	return fib(n-1) + fib(n-2)
//}

/*
带备忘录解法
1.时间复杂度O(n)，空间复杂度
*/
//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//
//	filterMap := map[int]int{}
//	filterMap[n+1] = 0
//
//	return helper(filterMap, n)
//
//}
//
//func helper(filterMap map[int]int, n int) int {
//
//	if n == 1 || n == 2 {
//		return 1
//	}
//
//	v, _ := filterMap[n]
//	if v != 0 {
//		return v
//	}
//
//	filterMap[n] = helper(filterMap, n-1) + helper(filterMap, n-2)
//	return filterMap[n]
//
//}

/*db数组的迭代解法*/
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}
	prev := 1
	curr := 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}

	return curr
}

func main() {
	res := fib(100)
	fmt.Println(res)

}
