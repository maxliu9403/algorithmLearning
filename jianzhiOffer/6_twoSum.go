package jianzhiOffer

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	p1, p2 := 0, n-1
	for p1 < p2 && numbers[p1]+numbers[p2] != target {
		if numbers[p1]+numbers[p2] < target {
			p1++
		} else {
			p2--
		}
	}
	ans := []int{p1, p2}
	return ans
}
