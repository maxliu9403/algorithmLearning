package isAnagram

import "sort"

/*
思路1：排序后判断
*/
func isAnagramSort(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	if string(s1) == string(s2) {
		return true
	}
	return false

}

/*
思路2：由于字母异位词，字母的出现个数是一致的，可以通过切片索引来记录每个字符的出现次数
时间复杂度：O(n)，其中 n 为 s 的长度。
空间复杂度：O(S)，其中 s 为字符集大小，此处 S=26。
*/
func isAnagramMaxLiu(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ss, ts [26]int
	for _, ch := range s {
		ss[ch-'a']++
	}
	for _, ch := range t {
		ts[ch-'a']++
	}
	if ss == ts {
		return true
	}
	return false
}
