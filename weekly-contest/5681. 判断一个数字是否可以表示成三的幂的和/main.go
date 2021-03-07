package main

import (
	"fmt"
)

//5681. 判断一个数字是否可以表示成三的幂的和
//https://leetcode-cn.com/contest/biweekly-contest-47/problems/check-if-number-is-a-sum-of-powers-of-three/
func checkPowersOfThree(n int) bool {
	//too young, too simple
	for n != 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	res := checkPowersOfThree(91)
	fmt.Print("res = ", res)
}
