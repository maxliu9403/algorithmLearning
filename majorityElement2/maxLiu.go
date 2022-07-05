package majorityElement2

func majorityElement(nums []int) (ans []int) {
	// 初始化两个候选人
	element1, element2 := 0, 0
	// 初始化两个候选人票数
	vote1, vote2 := 0, 0
	for _, num := range nums {
		// 为当前候选人
		if vote1 > 0 && element1 == num {
			vote1++
		} else if vote2 > 0 && element2 == num {
			vote2++
			// 第一个候选人票数为0
		} else if vote1 == 0 {
			element1 = num
			vote1++
		} else if vote2 == 0 {
			element2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	// 开始计数
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == element1 {
			cnt1++
		}
		if vote2 > 0 && num == element2 {
			cnt2++
		}
	}
	// 检测元素出现的次数是否满足要求
	if vote1 > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, element1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, element2)
	}
	return
}
