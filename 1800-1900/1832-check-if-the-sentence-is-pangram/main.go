package main

// https://leetcode.cn/problems/check-if-the-sentence-is-pangram/
// 1832. 判断句子是否为全字母句
func checkIfPangram(sentence string) bool {
	ans := make([]int, 26)
	for i := 0; i < len(sentence); i++ {
		ans[sentence[i]-'a']++
	}
	for _, an := range ans {
		if an == 0 {
			return false
		}
	}
	return true
}
func main() {

}
