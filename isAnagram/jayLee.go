package isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	f := map[rune]int{}
	for _, sv := range s {
		f[sv]++
	}
	for _, tv := range t {
		f[tv]--
		if f[tv] < 0 {
			return false
		}
	}
	return true
}
