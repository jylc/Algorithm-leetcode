package main

import "fmt"

//剑指 Offer 15. 二进制中1的个数
//https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	cnt := 0
	for i := 1; i <= 32; i++ {
		if num&1 > 0 {
			cnt++
		}
		num = num >> 1
	}
	return cnt
}
func main() {
	res := hammingWeight(00000000000000000000000000001011)
	fmt.Printf("result = %v\n", res)
}
