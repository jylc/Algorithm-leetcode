package main

import "fmt"

//https://leetcode-cn.com/problems/recursive-mulitply-lcci/
//面试题 08.05. 递归乘法
func minAndMax(A, B int) (min, max int) {
	if A < B {
		min = A
		max = B
	} else {
		min = B
		max = A
	}
	return
}

func multiply(A int, B int) int {
	min, max := minAndMax(A, B)
	ans := 0
	for i := 0; min != 0; i++ {
		if min&1 != 0 {
			ans += max << i
		}
		min >>= 1
	}
	return ans
}

func main() {
	ans := multiply(3, 4)
	//ans := 4 & 4
	fmt.Println(ans)
}
