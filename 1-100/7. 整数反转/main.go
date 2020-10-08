package main

import (
	"fmt"
	"math"
)

//7. 整数反转
//https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	tmp := int(math.Abs(float64(x)))
	res := 0
	for tmp != 0 {
		res = res * 10
		res += tmp % 10
		tmp = tmp / 10
	}

	if x < 0 {
		res = res * (-1)
	}
	if res <= math.MaxInt32 && res >= math.MinInt32 {
		return res
	} else {
		return 0
	}
}
func main() {
	fmt.Print("result: ", reverse(2147483647))
}
