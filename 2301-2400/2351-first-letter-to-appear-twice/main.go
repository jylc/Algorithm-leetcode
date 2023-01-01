package main

// https://leetcode.cn/problems/first-letter-to-appear-twice/
// 2351. 第一个出现两次的字母

func repeatedCharacter(s string) byte {
	m := make(map[byte]int)
	bytes := []byte(s)
	for _, b := range bytes {
		m[b]++
		if m[b] == 2 {
			return b
		}
	}
	return 0
}

func main() {

}
