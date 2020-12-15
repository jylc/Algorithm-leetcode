package main

import (
	"fmt"
	"strconv"
)

//738. 单调递增的数字
//https://leetcode-cn.com/problems/monotone-increasing-digits/
func monotoneIncreasingDigits(N int) int {
	s := []byte(strconv.Itoa(N))
	i := 1
	for i < len(s) && s[i-1] <= s[i] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
func main() {
	res := monotoneIncreasingDigits(9)
	fmt.Print(res)
}
