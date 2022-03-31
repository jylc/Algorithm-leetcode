package main

import (
	"fmt"
	"math/bits"
)

//https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/
//762. 二进制表示中质数个计算置位
func countPrimeSetBits(l int, r int) int {
	var res int
	for i := l; i <= r; i++ {
		ones := bits.OnesCount(uint(i))
		if check(ones) {
			res += 1
		}
	}
	return res
}

func check(ones int) bool {
	switch ones {
	case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41:
		return true
	}
	return false
}

func main() {
	ans := countPrimeSetBits(10, 15)
	fmt.Println(ans)
}
