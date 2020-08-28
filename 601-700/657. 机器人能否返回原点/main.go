package main

import (
	"fmt"
	"strings"
)

//657. 机器人能否返回原点
//https://leetcode-cn.com/problems/robot-return-to-origin/
func judgeCircle(moves string) bool {
	a, b := 0, 0
	for _, move := range moves {
		if strings.Compare(string(move), "U") == 0 {
			a++
		} else if strings.Compare(string(move), "R") == 0 {
			b++
		} else if strings.Compare(string(move), "L") == 0 {
			b--
		} else if strings.Compare(string(move), "D") == 0 {
			a--
		}
	}
	if a == 0 && b == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	result := judgeCircle("LL")
	fmt.Println("result = ", result)
}
