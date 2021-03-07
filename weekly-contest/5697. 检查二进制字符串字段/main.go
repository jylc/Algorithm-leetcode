package main

import "fmt"

//5697. 检查二进制字符串字段
//https://leetcode-cn.com/contest/weekly-contest-231/problems/check-if-binary-string-has-at-most-one-segment-of-ones/

/*
思路：从头开始遍历，遇到时0先标记，再次遇到1时，如果之前已经遇到0了则表明1不连续返回false（字符串由字符1起始）
*/
func checkOnesSegment(s string) bool {
	n := len(s)
	flag := false
	for i := 0; i < n; i++ {
		if flag && s[i] == '1' {
			return false
		}
		if s[i] == '0' {
			flag = true
		}
	}
	return true
}

/*
思路：两边开始遍历，使i，j分别指向连续1的起始和终止处，在i,j之间遍历，若有0则返回false
*/
func checkOnesSegment1(s string) bool {
	i, j := 0, len(s)-1
	for i <= j && s[i] == '0' {
		i++
	}
	for i <= j && s[j] == '0' {
		j--
	}

	for k := i; k < j; k++ {
		if s[i] == '0' {
			return false
		}
	}
	return true
}

func main() {
	s := "1110000"
	segment := checkOnesSegment(s)
	fmt.Print("res = ", segment)
}
