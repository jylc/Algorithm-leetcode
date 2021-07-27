package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits/
//1736. 替换隐藏数字得到的最晚时间
func maximumTime(time string) string {
	timeStr := []rune(time)
	if timeStr[0] == '?' {
		if timeStr[1] != '?' && timeStr[1] > '3' {
			timeStr[0] = '1'
		}
		if timeStr[1] != '?' && timeStr[1] <= '3' {
			timeStr[0] = '2'
		}
		if timeStr[1] == '?' {
			timeStr[0] = '2'
			timeStr[1] = '3'
		}
	}
	if timeStr[1] == '?' {
		if timeStr[0] < '2' {
			timeStr[1] = '9'
		} else {
			timeStr[1] = '3'
		}
	}

	if timeStr[3] == '?' {
		timeStr[3] = '5'
	}
	if timeStr[4] == '?' {
		timeStr[4] = '9'
	}
	return string(timeStr)
}
func main() {
	res := maximumTime("?0:15")
	fmt.Println(res)
}
