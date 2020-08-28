package main

import "fmt"

//191. 位1的个数
//https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}
func main() {
	result := hammingWeight(128)
	fmt.Println("result = ", result)
}
