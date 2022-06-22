package majorityElement

// hash
func majorityElementMaxLiu(nums []int) int {
	tmp := map[int]int{}
	n := len(nums) / 2
	for _, num := range nums {
		tmp[num] ++
		if tmp[num] > n {
			return num
		}
	}
	return 0
}

// 摩尔投票法（Boyer–Moore majority vote algorithm）
/*
	思路：算法分为两个阶段
		1. 对抗阶段：分属两个候选人的票数进行两两对抗抵消
		2. 计数阶段：计算对抗结果中最后留下的候选人票数是否有效
		https://leetcode.cn/problems/majority-element/solution/tu-jie-mo-er-tou-piao-fa-python-go-by-jalan/
	局限性：只有众数过半的情况的才能使用，其实就是把所有候选人抽象为2个，一个是最多票数的人，一个是没有票数的，如果最多票数的人小于n/2，那么通过两两对抗的时候最终的结果将是0
*/
func majorityElementBoyerMooreMajorityVoteAlgorithm(nums []int) int {
	// 记录当前票数最多的人
	major := 0
	// 记录当前票数
	count := 0
	for _, num := range nums {
		// 如果当前票数为0说明前面所有的选票都已经对冲完了，此时的候选人就是票数最多的人
		if count == 0 {
			major = num
		}
		// 如果当前票数最多的人就是此时的候选人，那么当前票数加1
		if major == num {
			count++
			// 如果当前票数最多的人不是此时的候选人，那么当前票数减1，减到0则换人
		} else {
			count--
		}
	}
	return major
}
