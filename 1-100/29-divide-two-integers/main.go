package main

import "math"

//https://leetcode.cn/problems/divide-two-integers/
//29. 两数相除
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	a := dividend
	b := divisor
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res := div(a, b)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	}
	return -res
}
func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tb := b
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	return count + div(a-tb, b)
}

func main() {

}
