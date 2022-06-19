package groupAnagrams

// GroupAnagrams
// 字母出现次数作为key(字母为小写)：[26]int -> [a-z]作为数组index，出现次数作为数组的value
// 注：golang禁止slice作为map key
func GroupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
