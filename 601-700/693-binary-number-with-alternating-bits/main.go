package main

import "fmt"

//https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
//693. 交替位二进制数
func hasAlternatingBits(n int) bool {
	res := -1
	for n != 0 {
		tmp := n % 2
		if tmp == res {
			return false
		} else {
			res = tmp
		}
		n /= 2
	}
	return true
}
func main() {
	res := hasAlternatingBits(11)
	fmt.Println(res)
}
