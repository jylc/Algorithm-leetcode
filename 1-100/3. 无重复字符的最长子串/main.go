package main

//3. 无重复字符的最长子串
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if s == " " {
		return 1
	}
	if len(s) == 1 {
		return 1
	}
	m := make(map[uint8]bool)
	tmp, max := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; !ok {
				m[s[j]] = true
				tmp++
			} else {
				if tmp > max {
					max = tmp
				}
				m = make(map[uint8]bool)
				tmp = 0
				break
			}
		}
	}
	return max
}
func main() {

}
