package sortColors

func sortColorsMaxLiu(nums []int) {
	// p0交换0 p1交换1
	var p0, p1 int
	for i, num := range nums {
		if num == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			// 此时的p0可能是1，和i交换之后会把1交换出去，所以需要再将1交换回来
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			// 交换1
			p0++
			p1++
		} else if num == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}
