package main

import (
	"fmt"
	"strings"
)

//https://leetcode.cn/problems/largest-3-same-digit-number-in-string/
//2264. 字符串中最大的 3 位相同数字
func largestGoodInteger1(num string) string {
	l, r := 0, 0
	var tmp uint8
	tmp = '0' - 1
	for r < len(num) {
		r = l + 1
		for r < len(num) && num[r] == num[l] && r-l <= 2 {
			r++
		}
		if r-l == 3 {
			if tmp < num[l] {
				tmp = num[l]
			}
		}
		l = r
	}
	if tmp == '0'-1 {
		return ""
	} else {
		ans := strings.Repeat(string(tmp), 3)
		return ans
	}
}
func largestGoodInteger(num string) string {
	for c := '9'; c >= '0'; c-- {
		s := strings.Repeat(string(c), 3)
		if strings.Contains(num, s) {
			return s
		}
	}
	return ""
}
func main() {
	num := "42352338"
	ans := largestGoodInteger(num)
	fmt.Println(ans)
}
