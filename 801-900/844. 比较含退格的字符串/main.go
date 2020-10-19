package main

import "strings"

//844. 比较含退格的字符串
//https://leetcode-cn.com/problems/backspace-string-compare/
func backspaceCompare(S string, T string) bool {
	a, b := make([]rune, 0), make([]rune, 0)
	for _, s := range S {
		if s != '#' {
			a = append(a, s)
		} else {
			if len(a) != 0 {
				a = a[:len(a)-1]
			}
		}
	}

	for _, t := range T {
		if t != '#' {
			b = append(b, t)
		} else {
			if len(b) != 0 {
				b = b[:len(b)-1]
			}
		}
	}

	if strings.Compare(string(a), string(b)) == 0 {
		return true
	} else {
		return false
	}
}
func main() {

}
