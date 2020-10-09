package main

//58. 最后一个单词的长度
//https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	length := len(s) - 1
	cnt := 0
	for length >= 0 {
		if s[length] != ' ' {
			cnt++
		} else {
			if cnt != 0 {
				return cnt
			}
		}
		length--
	}

	return cnt
}
func main() {

}
