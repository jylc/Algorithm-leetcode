package main

import "fmt"

//https://leetcode-cn.com/problems/xor-operation-in-an-array/
//1486. 数组异或操作
func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}
func main() {
	n := 10
	start := 5
	ans := xorOperation(n, start)
	fmt.Println(ans)
}
