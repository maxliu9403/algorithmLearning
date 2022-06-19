package groupAnagrams

import (
	"sort"
)

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次

/*
	解方1：排序
	字母异位说明字母是相同的，可以将字符串转换成[]byte类型，然后对[]byte进行排序，对其str转换为新的str，如果新的str相同，则是异位。最后将新的str作为map的键记录历史str
*/

/*
复杂度分析

	时间复杂度：O(nklogk)，其中n是strs中的字符串的数量，k是strs中的字符串的的最大长度。需要遍历n个字符串，对于每个字符串，需要O(klogk)的时间进行排序以及O(1)的时间更新哈希表，因此总时间复杂度是 O(nklogk)。

	空间复杂度：O(nk)，其中n是strs中的字符串的数量，k是strs中的字符串的的最大长度。需要用哈希表存储全部字符串。
*/
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		// 按照相同的顺序排序
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

/*
	解方2：计数
	复杂度分析：
		时间复杂度：O(len(strs)*(strs中最长str长度 + 字符集长度))，字符集=26（全是小写字母）+字符集长度的原因是更新切片的耗时，以及O(1)的时间更新hashMap
		空间复杂度: O(len(strs) * [(strs中最长str长度) + 字符集长度])，本题中字符集为所有小写字母

*/
func groupAnagrams2(strs []string) [][]string {
	// 构建一个使用数组为key的map
	mp := map[[26]int][]string{}
	for _, str := range strs {
		// 定义一个长度为26的切片记录每个字母出现的次数
		tmp := [26]int{}
		for _, b := range str {
			// 通过判断字符串和字符串a的差值，通过切片索引记录出现次数
			tmp[b-'a']++
		}
		// 异位字符串每个字符串出现的次数是一样的，所以tmp是一样的
		mp[tmp] = append(mp[tmp], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
