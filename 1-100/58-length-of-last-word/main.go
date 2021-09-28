package main

//https://leetcode-cn.com/problems/length-of-last-word/
//58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	cnt := 0
	for s[index] != ' ' && index >= 0 {
		cnt++
		index--
	}
	return cnt
}
func main() {

}
