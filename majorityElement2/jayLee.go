package majorityElement2

// MajorityElement 摩尔选举法。
// 1. 确定候选人数 <= 2, 因此设置2个候选人就够用了;
// 2. 如有遇到两个候选人之外的元素, 抵消2个候选人的count;
// nums:       [7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7]
// candidate1: [7, 7, _, 7, _, 7, _, 7, _, _, 7, 7, 7, 7, 7, 7]
// count1:     [1, 2, 2, 3, 3, 2, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9]
// candidate2: [_, _, 5, _, 5, 5, 5, _, 5, 5, _, _, _, _, _, _]
// count2:     [_, _, 1, 1, 2, 1, 2, _, 3, 4, 4, 4, 4, 4, 4, 4]
// 3. 判断是否满足众数要求:
//	[1, 2, 3, 4] => []，cd1 = 4, ct1 = 1, cd2 = 2, ct2 = 0
//	[1, 1, 2, 3] => [1], cd1 = 1, ct1 = 1, cd2 = 2, ct2 = 0
//	[1, 1, 2, 2] => [1, 2], cd1 = 1, ct1 = 2, cd2 = 2, ct2 = 2
// 举例说明，得到的count无法直接获得众数(参考case1的数组), 因此需要得到count>0的候选人的真实数量realCount.
func MajorityElement(nums []int) []int {
	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0
	for _, num := range nums {
		if count1 > 0 && candidate1 == num {
			count1++
		} else if count2 > 0 && candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1++
		} else if count2 == 0 {
			candidate2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}
	realCount1, realCount2 := 0, 0
	for _, num := range nums {
		if count1 > 0 && num == candidate1 {
			realCount1++
		}
		if count2 > 0 && num == candidate2 {
			realCount2++
		}
	}
	var response []int
	if realCount1 > len(nums)/3 {
		response = append(response, candidate1)
	}
	if realCount2 > len(nums)/3 {
		response = append(response, candidate2)
	}
	return response
}
