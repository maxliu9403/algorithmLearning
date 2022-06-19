package sortColors

func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	// 需要用到2个指针，P0 - 指向待替换为0的位置，P1 - 指向待替换为1的位置
	var p0, p1 int
	for idx, num := range nums {
		if num == 0 { // 值为0，交换 P0 与 idx，P0 P1各后移一位
			nums[idx], nums[p0] = nums[p0], nums[idx]
			if p0 < p1 { // 说明 p0 所在的位置的值为1；因此，1替换成0后，需要将1在替换到p1所在的位置
				nums[idx], nums[p1] = nums[p1], nums[idx]
			}
			p0, p1 = p0+1, p1+1
		} else if num == 1 { // 值为1，交换 P1 与 idx，P0位置不变，P1后移一位
			nums[idx], nums[p1] = nums[p1], nums[idx]
			p1++
		}
	}
}
