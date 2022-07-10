package topKFrequent

/*
	思路：桶排序，map记录元素出现的次数，创建一个数组，将频率作为数组下标，对于出现频率不同的数字集合，存入对应的数组下标
	时间复杂度：O(n);空间复杂度:O(n)
*/
func topKFrequent(nums []int, k int) []int {
	// 使用字典，统计每个元素出现的次数，元素为键，元素出现的次数为值
	ans := make([]int, 0)
	tmp := map[int]int{}
	for _, num := range nums {
		tmp[num] ++
	}
	// 创建桶
	// 初始长度len(nums)+1因为元素最大出现频率=len(nums)，索引从0开始，所以长度必须为len(nums)+1
	bucket := make([][]int, len(nums)+1)
	// 桶排序
	// 将频率作为数组下标，对于出现频率不同的数字集合，存入对应的数组下标
	for key, v := range tmp {
		bucket[v] = append(bucket[v], key)
	}
	// 倒序遍历桶
	// i >= 0 && len(ans) < k 控制只有前k个元素
	for i := len(bucket) - 1; i >= 0 && len(ans) < k; i-- {
		if bucket[i] == nil {
			continue
		}
		ans = append(ans, bucket[i]...)
	}
	return ans
}
