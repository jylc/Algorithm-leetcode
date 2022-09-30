package main

// https://leetcode.cn/problems/check-permutation-lcci/
// 面试题 01.02. 判定是否互为字符重排

func CheckPermutation(s1 string, s2 string) bool {
	strMap := make(map[rune]int)
	for _, v := range s1 {
		strMap[v]++
	}

	for _, v := range s2 {
		strMap[v]--
	}

	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {

}
